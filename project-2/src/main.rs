use std::io;
use std::collections::BTreeSet;
use std::cmp::Ordering;

// Manatee struct
#[derive(Debug)]
#[derive(Eq)]
#[derive(PartialEq)]
#[derive(PartialOrd)]
pub struct Manatee {
	size: i32,
	age: i32,
	index: i32
}
impl Ord for Manatee {
	fn cmp(&self, other: &Self) -> Ordering {
		if self.age == other.age {
			if self.size == other.size {
				return self.index.cmp(&other.index);
			}
			return self.size.cmp(&other.size);
		}
		return self.age.cmp(&other.age);
	}
}

// Main function
fn main() -> io::Result<()> {	
	let mut pairs: usize = 0;
	let mut female_manatees: BTreeSet<Manatee> = BTreeSet::new();
	let mut male_manatees: BTreeSet<Manatee> = BTreeSet::new();
	
	// Read the first five lines from stdio
	for n in 1..6  {
		// Set the input variable to be a new string
		let mut buffer = String::new();
		// Set the next input equal to the input variable
		io::stdin().read_line(&mut buffer)?;
		let mut buffer_split = buffer.split_whitespace();
		
		match n {
			// Get number of pairs
			1 =>
			pairs = buffer_split.next().expect("found empty string where number expected").parse::<usize>().unwrap(),
			
			// Get female manatee age
			2 => 
			for i in 1..(pairs+1) {
				let ind: i32 = i as i32;
				let mut empty_manatee = Manatee {
					size: 0,
					age: 0,
					index: ind
				};
				empty_manatee.age = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
				female_manatees.insert(empty_manatee);
				
			},
			
			// Get female manatee size
			3 => {
				let mut manatee_iter = female_manatees.iter();
				for i in 1..(pairs+1) {
					manatee_iter.next().unwrap().size = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
				}
			},
			
			// Get male manatee age
			4 =>
			for i in 1..(pairs+1) {
				let ind: i32 = i as i32;
				let mut empty_manatee = Manatee {
					size: 0,
					age: 0,
					index: ind
				};
				empty_manatee.age = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
				male_manatees.insert(empty_manatee);
			},
			
			// Get male manatee size
			5 => {
				let mut manatee_iter = male_manatees.iter();
				for i in 1..(pairs+1) {
					manatee_iter.next().unwrap().size = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
				}
			},
			
			// For anything else (which shouldn't occur)
			_ =>
			break,
		}			
	}


	print!("Female Manatees: \n");
	let mut manatee_iter = female_manatees.iter();
	for i in 1..(pairs+1) {
		let current_manatee = manatee_iter.next().unwrap();
		print!("S: {}  A: {}  I: {}\n", current_manatee.size, current_manatee.age, current_manatee.index);
	}
	Ok(())
}



// Unused sorting algorithm
/*for j in 1..(pairs - 1) {
	let mut swap_index = j;
	for n in j..pairs {
		if female_manatees[n].age <= female_manatees[swap_index].age {
			swap_index = n;
		}
	}
	let temp = female_manatees[j];
	female_manatees[j] = female_manatees[swap_index];
	female_manatees[swap_index] = temp;
}
// then the male ones
for j in 1..(pairs - 1) {
	let mut swap_index = j;
	for n in j..pairs {
		if &male_manatees[n].age <= &male_manatees[swap_index].age {
			swap_index = n;
		}
	}
	let temp = male_manatees[j];
	male_manatees[j] = male_manatees[swap_index];
	male_manatees[swap_index] = temp;
}
*/