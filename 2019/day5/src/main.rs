use anyhow::Result;
use log::{debug, info};
use thiserror::Error;

use std::{collections::HashMap, env, fs, str::FromStr};

fn main() -> Result<()> {
    env::set_var("RUST_LOG", "day5=trace");

    env_logger::init();

    let input = read_input()?;
    info!("{:?}", input);
    let output = run_intcode(input);
    info!("output = {:?}", output);
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

fn run_intcode(input: Vec<i64>) -> Vec<i64> {
    let mut input = input.clone();
    let mut last_result = ComputationResult::Computed { move_forward: 0 };
    let mut index = 0;
    let mut outputs = Vec::new();
    loop {
        last_result = execute(&mut input, index);
        match last_result {
            ComputationResult::Output {
                output,
                move_forward,
            } => {
                outputs.push(output);
                index += move_forward;
            }
            ComputationResult::Computed { move_forward } => {
                index += move_forward;
            }
            ComputationResult::Jump(new_index) => {
                index = new_index;
            }
            ComputationResult::Stop => {
                dbg!("Stopped");
                break;
            }
        }
    }
    outputs
}

fn execute(input: &mut Vec<i64>, index: usize) -> ComputationResult {
    let opcode = input[index];

    match opcode.to_string().parse::<OpCode>().unwrap() {
        OpCode::Add(parameter1, parameter2) => {
            let (arg1_i, arg2_i, result_i) = (input[index + 1], input[index + 2], input[index + 3]);
            let arg1 = if parameter1 == Parameter::PositionMode {
                input[arg1_i as usize]
            } else {
                arg1_i
            };
            let arg2 = if parameter2 == Parameter::PositionMode {
                input[arg2_i as usize]
            } else {
                arg2_i
            };
            let result = arg1 + arg2;
            input[result_i as usize] = result;

            ComputationResult::Computed { move_forward: 4 }
        }
        OpCode::Multiply(parameter1, parameter2) => {
            let (arg1_i, arg2_i, result_i) = (input[index + 1], input[index + 2], input[index + 3]);
            let arg1 = if parameter1 == Parameter::PositionMode {
                input[arg1_i as usize]
            } else {
                arg1_i
            };
            let arg2 = if parameter2 == Parameter::PositionMode {
                input[arg2_i as usize]
            } else {
                arg2_i
            };
            let result = arg1 * arg2;
            input[result_i as usize] = result;

            ComputationResult::Computed { move_forward: 4 }
        }
        OpCode::Stop => ComputationResult::Stop,
        OpCode::Input => {
            let arg = input[index + 1] as usize;
            input[arg] = 5;
            ComputationResult::Computed { move_forward: 2 }
        }
        OpCode::Output(parameter) => {
            let arg1_i = input[index + 1];
            let output = match parameter {
                Parameter::PositionMode => input[arg1_i as usize],
                Parameter::ImmediateMode => arg1_i,
            };
            ComputationResult::Output {
                move_forward: 2,
                output: output,
            }
        }
        OpCode::JumpIfTrue(parameter1, parameter2) => {
            let (arg1_i, arg2_i) = (input[index + 1], input[index + 2]);
            let arg1 = match parameter1 {
                Parameter::PositionMode => input[arg1_i as usize],
                Parameter::ImmediateMode => arg1_i,
            };
            if arg1 != 0 {
                let jump_index = match parameter2 {
                    Parameter::PositionMode => input[arg2_i as usize] as usize,
                    Parameter::ImmediateMode => arg2_i as usize,
                };
                ComputationResult::Jump(jump_index)
            } else {
                ComputationResult::Computed { move_forward: 3 }
            }
        }
        OpCode::JumpIfFalse(parameter1, parameter2) => {
            let (arg1_i, arg2_i) = (input[index + 1], input[index + 2]);
            let arg1 = match parameter1 {
                Parameter::PositionMode => input[arg1_i as usize],
                Parameter::ImmediateMode => arg1_i,
            };
            if arg1 == 0 {
                let jump_index = match parameter2 {
                    Parameter::PositionMode => input[arg2_i as usize] as usize,
                    Parameter::ImmediateMode => arg2_i as usize,
                };
                ComputationResult::Jump(jump_index)
            } else {
                ComputationResult::Computed { move_forward: 3 }
            }
        }
        OpCode::LessThan(parameter1, parameter2) => {
            let (arg1_i, arg2_i, result_i) = (input[index + 1], input[index + 2], input[index + 3]);
            let arg1 = match parameter1 {
                Parameter::PositionMode => input[arg1_i as usize],
                Parameter::ImmediateMode => arg1_i,
            };
            let arg2 = match parameter2 {
                Parameter::PositionMode => input[arg2_i as usize],
                Parameter::ImmediateMode => arg2_i,
            };
            if arg1 < arg2 {
                input[result_i as usize] = 1;
            } else {
                input[result_i as usize] = 0;
            }
            ComputationResult::Computed { move_forward: 4 }
        }
        OpCode::Equals(parameter1, parameter2) => {
            let (arg1_i, arg2_i, result_i) = (input[index + 1], input[index + 2], input[index + 3]);
            let arg1 = match parameter1 {
                Parameter::PositionMode => input[arg1_i as usize],
                Parameter::ImmediateMode => arg1_i,
            };
            let arg2 = match parameter2 {
                Parameter::PositionMode => input[arg2_i as usize],
                Parameter::ImmediateMode => arg2_i,
            };
            if arg1 == arg2 {
                input[result_i as usize] = 1;
            } else {
                input[result_i as usize] = 0;
            }
            ComputationResult::Computed { move_forward: 4 }
        }
    }
}

pub enum OpCode {
    Add(Parameter, Parameter),
    Multiply(Parameter, Parameter),
    Input,
    Output(Parameter),
    JumpIfTrue(Parameter, Parameter),
    JumpIfFalse(Parameter, Parameter),
    LessThan(Parameter, Parameter),
    Equals(Parameter, Parameter),
    Stop,
}

impl OpCode {
    fn from_i64(code: i64) -> OpCode {
        match code {
            1 => OpCode::Add(Parameter::PositionMode, Parameter::PositionMode),
            2 => OpCode::Multiply(Parameter::PositionMode, Parameter::PositionMode),
            3 => OpCode::Input,
            4 => OpCode::Output(Parameter::PositionMode),
            5 => OpCode::JumpIfTrue(Parameter::PositionMode, Parameter::PositionMode),
            6 => OpCode::JumpIfFalse(Parameter::PositionMode, Parameter::PositionMode),
            7 => OpCode::LessThan(Parameter::PositionMode, Parameter::PositionMode),
            8 => OpCode::Equals(Parameter::PositionMode, Parameter::PositionMode),
            99 => OpCode::Stop,
            _ => panic!("Unrecognized Opcode {}", code),
        }
    }
}

impl FromStr for OpCode {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        if s.len() <= 2 {
            Ok(OpCode::from_i64(s.parse::<i64>()?))
        } else {
            let opcode = s
                .chars()
                .rev()
                .take(2)
                .collect::<Vec<_>>()
                .into_iter()
                .rev()
                .collect::<String>()
                .parse::<i64>()?;
            let mut param1 = Parameter::PositionMode;
            let mut param2 = Parameter::PositionMode;

            s.chars().rev().skip(2).enumerate().for_each(|(i, x)| {
                if i == 0 && x == '1' {
                    param1 = Parameter::ImmediateMode;
                }
                if i == 1 && x == '1' {
                    param2 = Parameter::ImmediateMode;
                }
            });
            let parsed_opcode = OpCode::from_i64(opcode);
            let res_opcode = match parsed_opcode {
                OpCode::Add(_, _) => OpCode::Add(param1, param2),
                OpCode::Multiply(_, _) => OpCode::Multiply(param1, param2),
                OpCode::JumpIfTrue(_, _) => OpCode::JumpIfTrue(param1, param2),
                OpCode::JumpIfFalse(_, _) => OpCode::JumpIfFalse(param1, param2),
                OpCode::LessThan(_, _) => OpCode::LessThan(param1, param2),
                OpCode::Equals(_, _) => OpCode::Equals(param1, param2),
                OpCode::Output(_) => OpCode::Output(param1),
                _ => parsed_opcode,
            };
            Ok(res_opcode)
        }
    }
}

#[derive(PartialEq, Eq, Clone, Copy, Debug)]
pub enum Parameter {
    PositionMode,
    ImmediateMode,
}

#[derive(PartialEq, Eq, Clone, Copy, Debug)]
pub enum ComputationResult {
    Computed { move_forward: usize },
    Stop,
    Output { move_forward: usize, output: i64 },
    Jump(usize),
}
