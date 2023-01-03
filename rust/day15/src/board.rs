use regex::Regex;
use std::io::{BufRead, BufReader};
use std::{collections::HashSet, fs::File};

#[derive(PartialEq, Eq, Hash, Clone, Copy, Debug)]
struct Point {
    x: i32,
    y: i32,
}

#[derive(Debug)]
struct Measurement {
    sensor: Point,
    beacon: Point,
}

impl Measurement {
    pub fn new() -> Self {
        Measurement {
            sensor: Point { x: 0, y: 0 },
            beacon: Point { x: 0, y: 0 },
        }
    }

    pub fn parse_line(line: &str) -> Measurement {
        let pattern = Regex::new(r"([Ss]ensor|[Bb]eacon).*?x=(-?\d+),\s+y=(=?\d+)")
            .expect("Failed to compile Measurement Regex pattern.");
        let mut output = Measurement::new();
        let mut p: &mut Point;
        for capture in pattern.captures_iter(line) {
            if capture
                .get(1)
                .map_or(false, |m| m.as_str().to_lowercase() == "beacon")
            {
                // Beacon
                p = &mut output.beacon;
            } else if capture
                .get(1)
                .map_or(false, |m| m.as_str().to_lowercase() == "sensor")
            {
                // Sensor
                p = &mut output.sensor;
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
        return output;
    }
}

pub struct Board {
    data: Vec<Measurement>,
    sensors: HashSet<Point>,
    beacons: HashSet<Point>,
}

impl Board {
    pub fn parse_file(file_name: &str) -> Board {
        let mut output = Board {
            data: Vec::new(),
            sensors: HashSet::new(),
            beacons: HashSet::new(),
        };

        let fin = File::open(file_name).expect(&format!("Could not open {file_name}"));
        for line in BufReader::new(fin).lines() {
            output
                .data
                .push(Measurement::parse_line(line.as_ref().unwrap()));
            output
                .sensors
                .insert(output.data.last().expect("Missing sensor").sensor);
            output
                .beacons
                .insert(output.data.last().expect("Missing beacon").beacon);
        }

        return output;
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
        assert_eq!(b.data[0].sensor.x, 2);
        assert_eq!(b.data[0].sensor.y, 18);
        assert_eq!(b.data[0].beacon.x, -2);
        assert_eq!(b.data[0].beacon.y, 15);

        assert_eq!(b.data[3].sensor.x, 12);
        assert_eq!(b.data[3].sensor.y, 14);
        assert_eq!(b.data[3].beacon.x, 10);
        assert_eq!(b.data[3].beacon.y, 16);
    }
}
