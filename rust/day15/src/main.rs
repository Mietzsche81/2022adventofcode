use std::env;

mod board;

fn main() {
    // Read commandline arguments
    let args: Vec<String> = env::args().collect();
    if args.len() != 3 {
        panic!("Usage: ./day15 <1,2> <puzzle.input>")
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
    board::Board::parse_file(puzzle_input);
    println!("Hello 1!");
}

fn main2(_puzzle_input: &String) -> () {
    println!("Hello 2!");
}
