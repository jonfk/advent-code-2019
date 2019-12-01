use anyhow::Result;
use log::{info, trace};
use std::{env, fs};

fn main() -> Result<()> {
    env::set_var("RUST_LOG", "day1=trace");

    env_logger::init();

    let input = read_input()?;

    let total_fuel = input
        .into_iter()
        .map(|i| calculate_total_module_fuel(i))
        .fold(0, |sum, i| sum + i);

    info!("total fuel: {}", total_fuel);
    Ok(())
}

fn read_input() -> Result<Vec<i64>> {
    let input = fs::read_to_string("input.txt")?;

    let inputs: Result<Vec<_>> = input
        .lines()
        .filter(|line| !line.trim().is_empty())
        .map(|line| {
            let line = line.trim();
            Ok(line.parse::<i64>()?)
        })
        .collect();

    inputs
}

fn calculate_fuel(input: i64) -> i64 {
    (input / 3) - 2
}

fn calculate_total_module_fuel(input: i64) -> i64 {
    let mut current_fuel = calculate_fuel(input);
    let mut total_fuel = current_fuel;

    while current_fuel > 0 {
        current_fuel = calculate_fuel(current_fuel);
        if current_fuel > 0 {
            total_fuel = total_fuel + current_fuel;
        }
    }
    total_fuel
}
