use anyhow::Result;
use log::{debug, info};
use thiserror::Error;

use std::{collections::HashMap, env, fs, str::FromStr};

fn main() -> Result<()> {
    env::set_var("RUST_LOG", "day3=trace");

    env_logger::init();

    let input = read_input()?;

    let start = Loc::new(1, 1);
    let mut grid = Grid::new();
    debug!("created grid");

    for directions in input {
        grid.lay_wire(start, directions);
    }

    debug!("finding crossed wires");
    let crossed = grid.find_crossed();

    let mut closest: Option<Loc> = None;

    debug!("finding closest crossed wires");
    for loc in crossed {
        if let Some(cur_closest) = closest {
            if loc.distance(start) < cur_closest.distance(start) {
                closest = Some(loc);
            }
        } else {
            closest = Some(loc);
        }
    }

    info!("shortest distance: {}", closest.unwrap().distance(start));

    let cross_with_min_distance_travelled = grid.find_crossed_with_min_distance();

    info!(
        "shortest distance travelled cross: {:?}",
        cross_with_min_distance_travelled
    );

    Ok(())
}

fn read_input() -> Result<Vec<Vec<Direction>>> {
    let input = fs::read_to_string("input.txt")?;

    let inputs: Result<Vec<_>> = input
        .lines()
        .filter(|line| !line.trim().is_empty())
        .map(|line| {
            let directions = line.split(",");

            let directions: std::result::Result<Vec<Direction>, _> = directions
                .filter(|i| !i.trim().is_empty())
                .map(|d_str| d_str.parse::<Direction>())
                .collect();

            Ok(directions?)
        })
        .collect();

    inputs
}

#[derive(Copy, Clone, Debug)]
enum Direction {
    Up(i64),
    Down(i64),
    Right(i64),
    Left(i64),
}

impl Direction {
    fn moves(&self) -> i64 {
        match self {
            Direction::Up(moves) => *moves,
            Direction::Down(moves) => *moves,
            Direction::Right(moves) => *moves,
            Direction::Left(moves) => *moves,
        }
    }
}

#[derive(Error, Debug)]
pub enum DirectionParseError {
    #[error("Unknown direction")]
    Unknown,

    #[error("Couldn't parse number {}", .0)]
    ParseNumber(String),
}

impl FromStr for Direction {
    type Err = DirectionParseError;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        let (direction, num_str) = s.split_at(1);
        let num: i64 = num_str
            .parse()
            .map_err(|_e| DirectionParseError::ParseNumber(num_str.to_owned()))?;
        let dir = match direction {
            "U" => Direction::Up(num),
            "D" => Direction::Down(num),
            "R" => Direction::Right(num),
            "L" => Direction::Left(num),
            _ => return Err(DirectionParseError::Unknown),
        };
        Ok(dir)
    }
}

#[derive(Copy, Clone, Debug, PartialEq, Eq, Hash)]
struct Loc {
    pub x: i64,
    pub y: i64,
}

impl Loc {
    fn new(x: i64, y: i64) -> Loc {
        Loc { x: x, y: y }
    }

    fn distance(&self, other: Loc) -> i64 {
        (self.x - other.x).abs() + (self.y - other.y).abs()
    }
}

struct Cell {
    distance: i64,
    state: State,
}

impl Cell {
    fn new() -> Cell {
        Cell {
            distance: 0,
            state: State::Empty,
        }
    }
}

#[derive(Copy, Clone, Debug, PartialEq, Eq)]
enum State {
    Crossed,
    Occupied,
    Empty,
}

struct Grid {
    grid: HashMap<Loc, Cell>,
}

impl Grid {
    fn new() -> Grid {
        Grid {
            grid: HashMap::new(),
        }
    }

    fn lay_wire(&mut self, start: Loc, directions: Vec<Direction>) {
        let mut current = start;
        let mut distance = 0;
        self.grid
            .entry(Loc::new(start.x, start.y))
            .or_insert(Cell::new());

        for direction in directions {
            current = self.lay_direction(current, direction, distance);
            distance += direction.moves();
        }
    }

    fn lay_direction(&mut self, start: Loc, direction: Direction, distance: i64) -> Loc {
        match direction {
            Direction::Up(moves) => {
                for y in 1..=moves {
                    let entry = self
                        .grid
                        .entry(Loc::new(start.x, start.y + y))
                        .or_insert(Cell::new());
                    if entry.state == State::Empty {
                        entry.state = State::Occupied;
                    } else {
                        entry.state = State::Crossed;
                    }
                    entry.distance += distance + y;
                }
                Loc::new(start.x, start.y + moves)
            }
            Direction::Down(moves) => {
                for y in 1..=moves {
                    let entry = self
                        .grid
                        .entry(Loc::new(start.x, start.y - y))
                        .or_insert(Cell::new());
                    if entry.state == State::Empty {
                        entry.state = State::Occupied;
                    } else {
                        entry.state = State::Crossed;
                    }
                    entry.distance += distance + y;
                }
                Loc::new(start.x, start.y - moves)
            }
            Direction::Right(moves) => {
                for x in 1..=moves {
                    let entry = self
                        .grid
                        .entry(Loc::new(start.x + x, start.y))
                        .or_insert(Cell::new());
                    if entry.state == State::Empty {
                        entry.state = State::Occupied;
                    } else {
                        entry.state = State::Crossed;
                    }
                    entry.distance += distance + x;
                }
                Loc::new(start.x + moves, start.y)
            }
            Direction::Left(moves) => {
                for x in 1..=moves {
                    let entry = self
                        .grid
                        .entry(Loc::new(start.x - x, start.y))
                        .or_insert(Cell::new());
                    if entry.state == State::Empty {
                        entry.state = State::Occupied;
                    } else {
                        entry.state = State::Crossed;
                    }
                    entry.distance += distance + x;
                }
                Loc::new(start.x - moves, start.y)
            }
        }
    }

    fn find_crossed(&self) -> Vec<Loc> {
        self.grid
            .iter()
            .filter_map(|entry| {
                if entry.1.state == State::Crossed {
                    Some(*entry.0)
                } else {
                    None
                }
            })
            .collect()
    }

    fn find_crossed_with_min_distance(&self) -> Option<i64> {
        self.grid
            .iter()
            .filter_map(|entry| {
                if entry.1.state == State::Crossed && entry.0.x != 1 && entry.0.y != 1 {
                    Some(entry)
                } else {
                    None
                }
            })
            .min_by_key(|entry| entry.1.distance)
            .map(|entry| entry.1.distance)
    }
}
