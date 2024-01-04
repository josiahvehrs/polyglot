#![allow(clippy::needless_return)]

fn get_input() -> &'static str {
    "forward 5
down 5
forward 8
up 3
down 8
forward 2"
}

#[derive(Debug)]
struct Point {
    x: i32,
    y: i32,
}

fn parse_line(line: &'static str) -> Point {
    let (dir, amt) = line.split_once(' ').expect("should contain whitespace");
    let amount: i32 = str::parse(amt).expect("should be int");
    if dir == "forward" {
        return Point { x: amount, y: 0 };
    } else if dir == "up" {
        return Point { x: 0, y: -amount };
    }
    return Point { x: 0, y: amount };
}

fn main() {
    let result =
        get_input()
            .lines()
            .map(parse_line)
            .fold(Point { x: 0, y: 0 }, |mut acc, point| {
                acc.x += point.x;
                acc.y += point.y;
                return acc;
            });
    println!("{:?}", result)
}
