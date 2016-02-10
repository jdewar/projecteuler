// Almost heavily copied from http://rustbyexample.com/fn.html
fn main() {

// arrays are known at compile time, must supply [type; size].
    let divisors: [u32; 2] = [3, 5];
    let max_number = 1000;

    let mut accumulator = 0;

// `n` will take the values: 1, 2, ..., max_number-1 in each iteration
    for n in 1..max_number {
        let mut is_divisible = false;
        for i in 0..divisors.len() {
            let divisor = divisors[i];
            is_divisible = is_divisible_by(n,divisor);
            if is_divisible {
		break;
            }
        }
        if is_divisible {
            accumulator += n;
        }
    }
    println!("Total: {}",accumulator);
}

// Function that returns a boolean value
fn is_divisible_by(lhs: u32, rhs: u32) -> bool {
    // Corner case, early return
    if rhs == 0 {
        return false;
    }

    // This is an expression, the `return` keyword is not necessary here
    lhs % rhs == 0
}

