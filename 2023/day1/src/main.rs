use std::fs;

fn main() {
    let mut input_str = fs::read_to_string("day1/input.txt").unwrap();
    input_str = input_str.replace("oneight", "18");
    input_str = input_str.replace("threeight", "38");
    input_str = input_str.replace("fiveight", "58");
    input_str = input_str.replace("nineight", "98");
    input_str = input_str.replace("sevenine", "79");
    input_str = input_str.replace("eighthree", "83");
    input_str = input_str.replace("twone", "21");
    input_str = input_str.replace("eightwo", "82");
    input_str = input_str.replace("eighthree", "83");
    input_str = input_str.replace("one", "1");
    input_str = input_str.replace("two", "2");
    input_str = input_str.replace("three", "3");
    input_str = input_str.replace("four", "4");
    input_str = input_str.replace("five", "5");
    input_str = input_str.replace("six", "6");
    input_str = input_str.replace("seven", "7");
    input_str = input_str.replace("eight", "8");
    input_str = input_str.replace("nine", "9");

    let input: Vec<_> = input_str
        .split("\n")
        // .filter(|s| !s.trim().is_empty())
        .collect();

    // println!("input: {:?}", input_str);

    let mut sum = 0;

    for line in input {
        let digits: Vec<_> = line.chars().filter_map(|c| c.to_digit(10)).collect();

        if digits.len() < 1 {
            panic!("input line with no digits {}", line);
        }
        // println!("digits: {:?}", digits);

        let first = digits.first().unwrap();
        let second = digits.last().unwrap();
        sum += first * 10;
        sum += second;
    }
    println!("sum: {}", sum);
}
