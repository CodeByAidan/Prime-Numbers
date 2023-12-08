use std::fs::File;
use std::io::{self, Write};

fn is_prime(num: u32) -> bool {
    if num < 2 {
        return false;
    }

    (2..=(num as f64).sqrt() as u32).all(|i| num % i != 0)
}

fn main() {
    let mut file = File::create("check_prime.rs").expect("Error creating file");

    writeln!(file, "fn is_prime(num: u32) -> bool {{").unwrap();
    writeln!(file, "    if num < 2 {{").unwrap();
    writeln!(file, "        return false;").unwrap();
    writeln!(file, "    }}").unwrap();
    writeln!(file, "    (2..=(num as f64).sqrt() as u32).all(|i| num % i != 0)").unwrap();
    writeln!(file, "}}").unwrap();

    writeln!(file, "fn main() {{").unwrap();

    writeln!(file, "    println!(\"Enter a number:\");").unwrap();
    writeln!(file, "    let mut input = String::new();").unwrap();
    writeln!(file, "    std::io::stdin().read_line(&mut input).expect(\"Failed to read line\");").unwrap();
    writeln!(file, "    let num: u32 = input.trim().parse().expect(\"Please enter a valid number\");").unwrap();

    // make for loop to check all numbers from 1 to 1000
    for i in 1..=1000 {
        writeln!(file, "    if num == {} {{", i).unwrap();
        if is_prime(i) {
            writeln!(file, "        println!(\"{} is prime!\");", i).unwrap();
        } else {
            writeln!(file, "        println!(\"{} is not prime!\");", i).unwrap();
        }
        writeln!(file, "    }}").unwrap();
    }

    writeln!(file, "}}").unwrap();
}
