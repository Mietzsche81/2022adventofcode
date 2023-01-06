use std::env;
use std::fmt::Error;
use std::fs::File;
use std::io::{BufRead, BufReader};

mod snafu;
use crate::snafu::Conversion;

fn main() {
    // Read commandline arguments
    let args: Vec<String> = env::args().collect();
    if args.len() != 3 {
        panic!("Usage: cargo run <1,2> <puzzle.input>")
    }
    let puzzle_part = &args[1];
    let puzzle_input = &args[2];
    match puzzle_part.as_str() {
        "1" => main1(puzzle_input),
        "2" => main2(puzzle_input),
        _ => panic!("Must declare whether to execute part 1 or 2"),
    }
}

fn main1(puzzle_input: &String) -> () {
    let data = parse(puzzle_input).expect("Failed to collect data");
    let mut total = snafu::Snafu::new(0);
    for entry in data.iter() {
        total = total + entry;
    }
    println!("{:?} {:?}", total.to_base10(), total)
}

fn main2(puzzle_input: &String) -> () {}

fn parse(file_name: &String) -> Result<Vec<snafu::Snafu>, Error> {
    let fin = File::open(file_name).expect("Could not open file");
    let parser = BufReader::new(fin);
    let v = parser
        .lines()
        .map(|line| snafu::Snafu::parse(line.expect("Could not parse line").trim().to_string()))
        .collect();
    return Ok(v);
}
