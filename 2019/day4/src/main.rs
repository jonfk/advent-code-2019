fn main() {
    //dbg!(matches_criteria_part2(111144));
    println!(
        "number of matching passwords: {}",
        find_matches(273025, 767253)
    )
}

fn find_matches(start: i64, end: i64) -> usize {
    (start..=end).filter(|x| matches_criteria_part2(*x)).count()
}

fn matches_criteria_part1(password: i64) -> bool {
    let pass: Vec<_> = password
        .to_string()
        .chars()
        .map(|c| c.to_string().parse::<i64>().unwrap())
        .collect();

    let mut dedup_pass = pass.clone();
    dedup_pass.dedup();

    pass.len() == 6
        && dedup_pass.len() < 6
        && pass
            .into_iter()
            .fold(
                (true, 0),
                |acc, x| if acc.0 { ((acc.1 <= x), x) } else { acc },
            )
            .0
}

#[derive(Debug, Clone, Copy)]
struct Acc {
    matched: bool,
    x: i64,
    repetitions: usize,
}

impl Acc {
    fn new(matched: bool, x: i64, repetitions: usize) -> Acc {
        Acc {
            matched: matched,
            x: x,
            repetitions: repetitions,
        }
    }
}

fn matches_criteria_part2(password: i64) -> bool {
    let pass: Vec<_> = password
        .to_string()
        .chars()
        .map(|c| c.to_string().parse::<i64>().unwrap())
        .collect();

    let mut dedup_pass = pass.clone();
    dedup_pass.dedup();

    let only_double = pass
        .iter()
        .enumerate()
        .fold(
            Acc {
                matched: false,
                x: -1,
                repetitions: 0,
            },
            |acc, x| {
                if acc.matched {
                    acc
                } else {
                    if ((acc.x != *x.1) && acc.repetitions == 2)
                        || (x.0 == pass.len() - 1 && acc.repetitions == 1 && *x.1 == acc.x)
                    {
                        Acc::new(true, -1, 0)
                    } else if acc.x == *x.1 {
                        Acc::new(false, *x.1, acc.repetitions + 1)
                    } else {
                        Acc::new(false, *x.1, 1)
                    }
                }
            },
        )
        .matched;
    pass.len() == 6
        && only_double
        && pass
            .into_iter()
            .fold(
                (true, 0),
                |acc, x| if acc.0 { ((acc.1 <= x), x) } else { acc },
            )
            .0
}
