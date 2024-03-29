use std::{fs, u32};

//I'm new to Rust, sorry.

const NUMS: [&str; 10] = [
    "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
];

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();
    let mut sum1: u32 = 0;
    let mut sum2 = 0;

    for val in input.split('\n').into_iter() {
        if val.len() < 1 {
            continue;
        }
        let first;
        let last;
        let (first_digit, first_digit_index) = get_first_digit(val);
        let (last_digit, last_digit_index) = get_last_digit(val);
        sum1 += ((first_digit as u32) * 10) + (last_digit as u32);
        if let Some(f) = get_digit_by_text(&val[..(first_digit_index as usize)], true) {
            first = f as i32;
        } else {
            first = first_digit as i32;
        }
        if let Some(f) = get_digit_by_text(&val[(last_digit_index as usize)..], false) {
            last = f as i32;
        } else {
            last = last_digit as i32;
        }
        sum2 += (first * 10) + last;
    }
    println!("Part1: {sum1}");
    println!("Part2: {sum2}");
}

fn get_digit_by_text(s: &str, first: bool) -> Option<u8> {
    let mut index: Option<u8> = None;
    let mut number: Option<u8> = None;
    for (i, num) in NUMS.iter().enumerate() {
        if first {
            if let Some(num_index) = s.find(num) {
                if (num_index as u8) < index.unwrap_or(255) {
                    index = Some(num_index as u8);
                    number = Some(i as u8);
                }
            }
        } else {
            if let Some(num_index) = s.rfind(num) {
                if (num_index as u8) > index.unwrap_or(0) {
                    index = Some(num_index as u8);
                    number = Some(i as u8);
                }
            }
        }
    }
    number
}

fn get_first_digit(val: &str) -> (u8, u8) {
    for (i, v) in val.chars().enumerate() {
        if v.is_numeric() {
            return (v.to_digit(10).unwrap() as u8, i as u8);
        }
    }
    (255, 255)
}
fn get_last_digit(val: &str) -> (u8, u8) {
    for (i, v) in val.chars().rev().enumerate() {
        if v.is_numeric() {
            return (
                v.to_digit(10).unwrap() as u8,
                (val.len() as u8) - 1 - i as u8,
            );
        }
    }
    (0, 0)
}
