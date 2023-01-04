use regex::Regex;
use std::collections::HashMap;
use std::io::{BufRead, BufReader};
use std::ops::Sub;
use std::{collections::HashSet, fs::File};

#[derive(PartialEq, Eq, Hash, Clone, Copy, Debug)]
pub struct Point {
    pub x: i32,
    pub y: i32,
}

impl Sub<Point> for Point {
    type Output = i32;

    fn sub(self, other: Point) -> i32 {
        return (self.x - other.x).abs() + (self.y - other.y).abs();
    }
}

impl Sub<&Point> for &Point {
    type Output = i32;

    fn sub(self, other: &Point) -> i32 {
        return (self.x - other.x).abs() + (self.y - other.y).abs();
    }
}

#[derive(Clone, Copy, Debug)]
struct Measurement {
    sensor: Point,
    beacon: Point,
    distance: i32,
}

impl Measurement {
    pub fn new(sensor: Point, beacon: Point) -> Self {
        Measurement {
            sensor: sensor,
            beacon: beacon,
            distance: beacon - sensor,
        }
    }

    pub fn parse_line(line: &str) -> Measurement {
        let pattern = Regex::new(r"([Ss]ensor|[Bb]eacon).*?x=(-?\d+),\s+y=(=?\d+)")
            .expect("Failed to compile Measurement Regex pattern.");
        let mut beacon: Point = Point { x: 0, y: 0 };
        let mut sensor: Point = Point { x: 0, y: 0 };
        let mut p: &mut Point;
        for capture in pattern.captures_iter(line) {
            if capture
                .get(1)
                .map_or(false, |m| m.as_str().to_lowercase() == "beacon")
            {
                // Beacon
                p = &mut beacon;
            } else if capture
                .get(1)
                .map_or(false, |m| m.as_str().to_lowercase() == "sensor")
            {
                // Sensor
                p = &mut sensor;
            } else {
                panic!("Unable to parse line '{}'", line);
            }
            p.x = capture
                .get(2)
                .expect("Missing x coordinate")
                .as_str()
                .parse::<i32>()
                .expect("Could not parse x coordinate as int");
            p.y = capture
                .get(3)
                .expect("Missing y coordinate from beacon")
                .as_str()
                .parse::<i32>()
                .expect("Could not parse y coordinate as int");
        }
        return Measurement::new(sensor, beacon);
    }
}

#[derive(Clone, Debug)]
pub struct Board {
    sensors: HashMap<Point, Measurement>,
    beacons: HashSet<Point>,
    pub x_bound: (i32, i32),
    pub y_bound: (i32, i32),
}

impl Board {
    pub fn parse_file(file_name: &str) -> Board {
        let mut board = Board {
            sensors: HashMap::new(),
            beacons: HashSet::new(),
            x_bound: (0, 0),
            y_bound: (0, 0),
        };

        let fin = File::open(file_name).expect(&format!("Could not open {file_name}"));
        for line in BufReader::new(fin).lines() {
            let line = Measurement::parse_line(line.as_ref().unwrap());
            board.sensors.insert(line.sensor, line);
            board.beacons.insert(line.beacon);
        }
        board.determine_bounds();
        return board;
    }

    pub fn determine_bounds(&mut self) -> () {
        let mut x: Vec<i32> = Vec::new();
        let mut y: Vec<i32> = Vec::new();
        for m in self.sensors.values() {
            x.push(m.sensor.x - m.distance);
            x.push(m.sensor.x + m.distance);
            y.push(m.sensor.y - m.distance);
            y.push(m.sensor.y + m.distance);
        }
        self.x_bound = (
            *x.iter()
                .min()
                .expect("Cannot determine board X-bounds without sensor measurements"),
            *x.iter()
                .max()
                .expect("Cannot determine board X-bounds without sensor measurements"),
        );
        self.y_bound = (
            *y.iter()
                .min()
                .expect("Cannot determine board Y-bounds without sensor measurements"),
            *y.iter()
                .max()
                .expect("Cannot determine board Y-bounds without sensor measurements"),
        );
    }

    pub fn impossible_beacon(&self, query: &Point) -> bool {
        if self.beacons.contains(query) {
            return false;
        };

        for (sensor, measurement) in self.sensors.iter() {
            if (query - sensor) <= measurement.distance {
                return true;
            };
        }

        return false;
    }

    pub fn count_impossible_in_row(&self, y: i32) -> i32 {
        let mut sum: i32 = 0;

        for x in self.x_bound.0..=self.x_bound.1 {
            if self.impossible_beacon(&Point { x: x, y: y }) {
                sum += 1;
            }
        }
        return sum;
    }

    /// With a proper board/bounds, returns the only unscanned point
    ///
    /// The problem prescribes that there is only one valid point in board's
    /// bounds. Therefore, the point must be at some measurement's distance+1.
    /// Check each point around the board of each measurement's range, and stop when
    /// no other measurement interects (check via point < measurement.distance)
    pub fn find_only_empty(&self) -> Result<Point, ()> {
        let measurements = self.sensors.values().collect::<Vec<&Measurement>>();
        for m in &measurements {
            let r0 = [
                (m.sensor.x, m.sensor.y - m.distance - 1),
                (m.sensor.x + m.distance + 1, m.sensor.y),
                (m.sensor.x, m.sensor.y + m.distance + 1),
                (m.sensor.x - m.distance - 1, m.sensor.y),
            ];
            let dr = [(1, 1), (-1, 1), (-1, -1), (1, -1)];
            for i in 0..r0.len() {
                let mut p = Point {
                    x: r0[i].0,
                    y: r0[i].1,
                };
                if p.x < self.x_bound.0
                    || p.x > self.x_bound.1
                    || p.y < self.y_bound.0
                    || p.y > self.y_bound.1
                {
                    continue;
                }
                for _ in 0..=m.distance {
                    // Check
                    let mut valid = true;
                    for n in &measurements {
                        if (p - n.sensor) <= n.distance {
                            valid = false;
                            break;
                        }
                    }
                    if valid {
                        // We did it!
                        return Ok(p);
                    } else {
                        // Iterate
                        p.x += dr[i].0;
                        p.y += dr[i].1;
                    }
                }
            }
        }
        return Err(());
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn point_eq() {
        let a = Point { x: 4, y: 5 };
        let b = Point { x: 4, y: 5 };
        assert_eq!(a, b);
        assert_eq!(Point { y: -1, x: 100 }, Point { y: -1, x: 100 });
        // Just toying around with memory locations in rust...
        assert_eq!(&a as *const _, &a as *const _);
        assert_ne!(&a as *const _, &b as *const _);
    }

    #[test]
    fn board_parse() {
        let b = Board::parse_file("test.input");
        assert_eq!(b.beacons.len(), 6);
        assert_eq!(b.sensors.len(), 14);
        assert_eq!(b.x_bound, (-8, 28));
        assert_eq!(b.y_bound, (-10, 26));

        assert!(b.sensors.contains_key(&Point { x: 10, y: 20 }));
        assert!(!b.sensors.contains_key(&Point { x: 15, y: 3 }));
        assert_eq!(
            b.sensors[&Point { x: 16, y: 7 }].beacon,
            Point { x: 15, y: 3 }
        );
        assert_ne!(
            b.sensors[&Point { x: 17, y: 20 }].beacon,
            Point { x: 25, y: 17 }
        );

        assert!(b.beacons.contains(&Point { x: -2, y: 15 }));
        assert!(!b.beacons.contains(&Point { x: 2, y: 18 }));
    }

    #[test]
    fn board_count() {
        let b = Board::parse_file("test.input");
        assert_eq!(b.count_impossible_in_row(10), 26);
    }
}
