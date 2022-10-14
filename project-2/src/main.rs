use std::io;
use std::collections::BTreeSet;
use std::cmp::Ordering;

// Manatee struct
#[derive(Debug, Eq, PartialEq, PartialOrd, Copy, Clone)]
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
	
	let mut female_manatee_vec: Vec<Manatee> = Vec::new();
	let mut male_manatee_vec: Vec<Manatee> = Vec::new();
	
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
			
			// Get female manatee age and add it to the female_manatee_vec
			2 => 
			{
				for i in 1..(pairs+1) {
					let ind: i32 = i as i32;
					let mut empty_manatee = Manatee {
						size: 0,
						age: 0,
						index: ind
					};
					empty_manatee.age = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
					female_manatee_vec.push(empty_manatee);
				}
			},
			
			// Get female manatee size
			3 => 
			{
				for manatee in &mut female_manatee_vec {
					manatee.size = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
				}
			},
			
			// Get female manatee age and add it to the male_manatee_vec
			4 =>
			{
				for i in 1..(pairs+1) {
					let ind: i32 = i as i32;
					let mut empty_manatee = Manatee {
						size: 0,
						age: 0,
						index: ind
					};
					empty_manatee.age = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
					male_manatee_vec.push(empty_manatee);
				}
			},
			
			// Get male manatee size
			5 => {
				for manatee in &mut male_manatee_vec {
					manatee.size = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
				}
			},

			// For anything else (which shouldn't occur)
			_ => break,
		}			
	}

	// Insert all vector items into the BTrees
	for manatee in &mut male_manatee_vec {
	    male_manatees.insert(*manatee);
	}
	for manatee in &mut female_manatee_vec {
	    female_manatees.insert(*manatee);
	}

	// Print the manatee sets
	print!("Female Manatees: \n");
	let mut iter = female_manatees.iter();
	for _i in 1..(pairs+1) {
		let current_manatee = iter.next().unwrap();
		print!("S: {}  A: {}  I: {}\n", current_manatee.size, current_manatee.age, current_manatee.index);
	}
	print!("Male Manatees: \n");
	let mut iter = male_manatees.iter();
	for _i in 1..(pairs+1) {
		let current_manatee = iter.next().unwrap();
		print!("S: {}  A: {}  I: {}\n", current_manatee.size, current_manatee.age, current_manatee.index);
	}
	Ok(())
}