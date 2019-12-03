use anyhow::Result;
use log::{info, trace};
use std::{env, fs};

fn main() -> Result<()> {
    env::set_var("RUST_LOG", "day2=trace");

    env_logger::init();

    let input = read_input()?;
    let (noun, verb) = find_noun_verb(input, 19690720);

    info!(
        "noun: {}, verb: {}, result: 100*noun+verb = {}",
        noun,
        verb,
        (100 * noun + verb)
    );
    Ok(())
}

fn read_input() -> Result<Vec<i64>> {
    let input = fs::read_to_string("input.txt")?;

    let inputs: Result<Vec<_>> = input
        .split(",")
        .filter(|i| !i.trim().is_empty())
        .map(|i| {
            let i = i.trim();
            Ok(i.parse::<i64>()?)
        })
        .collect();

    inputs
}

fn find_noun_verb(input: Vec<i64>, expected: i64) -> (i64, i64) {
    for noun in 0..100 {
        for verb in 0..100 {
            let mut input = input.clone();
            update_intcode(&mut input, noun, verb);
            let result = run_intcode(input);
            if result == 19690720 {
                return (noun, verb);
            }
        }
    }
    (0, 0)
}

fn run_intcode(input: Vec<i64>) -> i64 {
    let mut input = input.clone();
    let mut last_result = ComputationResult::Computed;
    let mut index = 0;
    while last_result != ComputationResult::Stop && (index + 4) < input.len() {
        last_result = execute(&mut input, index);
        index += 4;
    }
    input[0]
}

fn update_intcode(input: &mut Vec<i64>, noun: i64, verb: i64) {
    input[1] = noun;
    input[2] = verb;
}

fn execute(input: &mut Vec<i64>, index: usize) -> ComputationResult {
    let opcode = input[index];
    let (arg1_i, arg2_i, result_i) = (input[index + 1], input[index + 2], input[index + 3]);

    match opcode {
        1 => {
            let arg1 = input[arg1_i as usize];
            let arg2 = input[arg2_i as usize];
            let result = arg1 + arg2;
            input[result_i as usize] = result;

            ComputationResult::Computed
        }
        2 => {
            let arg1 = input[arg1_i as usize];
            let arg2 = input[arg2_i as usize];
            let result = arg1 * arg2;
            input[result_i as usize] = result;

            ComputationResult::Computed
        }
        99 => ComputationResult::Stop,
        _ => unreachable!("unrecoginzed opcode"),
    }
}

#[derive(PartialEq, Eq, Clone, Copy, Debug)]
pub enum ComputationResult {
    Computed,
    Stop,
}
