use std::io;
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
	let mut female_manatees: Vec<Manatee> = Vec::new();
	let mut male_manatees: Vec<Manatee> = Vec::new();
	// let mut max_female_size: i32 = 0; // Seems unused?
	// let mut max_male_size: i32 = 0; // Also seems unused
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
					female_manatees.push(empty_manatee);
				}
			},
			
			// Get female manatee size
			3 => 
			{
				for manatee in &mut female_manatees {
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
					male_manatees.push(empty_manatee);
				}
			},
			
			// Get male manatee size
			5 => {
				for manatee in &mut male_manatees {
					manatee.size = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
				}
			},

			// For anything else (which shouldn't occur)
			_ => break,
		}			
	}


	
	// Sort both vectors of manatees.
	
	let mut female_answer: Vec<i32> = Vec::new();
	let mut male_answer: Vec<i32> = Vec::new();
	
	// morbin' time.

	// Assume no pair found until one is found
	let mut impossible = false;
	let mut no_pair = true;

	for i in 0..pairs-1 {
		// Sort both lists of manatees
		female_manatees.sort_by(|a, b| a.cmp(b));
		male_manatees.sort_by(|a, b| a.cmp(b));
		// Get the first manatee in each list
		let first_male = &male_manatees[i];
		let first_female = &female_manatees[i];
		// Get the smallest (first) female that is larger than the male.
		let youngest = first_female.age;
		for j in i..pairs {
			if(female_manatees[j].size > first_male.size) {
			   	print!("{}:{}", female_manatees[j].age, female_manatees[i].size); // Debug line; remove later
				female_answer.push(female_manatees[j].index);
				male_answer.push(male_manatees[i].index);
				no_pair = false;
				break;
			}
			// If you've passed the correct manatee age range, break off
			if(female_manatees[j].age > youngest) {
				break;
			}
		}
		if(no_pair) {
			// No pair? This manatee might match up with an older manatee.
			// Find the best match for the female, instead.
            let mut best_match = first_male;
			let youngest = first_male.age;
			// Finding the best male for the first female is a bit more complex.
			// We want the LARGEST one, but we don't know which is the largest
			// until we've polled them all (of the right age).
			for j in i..pairs {
				if((best_match.size < male_manatees[j].size && male_manatees[j].size < first_female.size) || best_match.size > first_female.size) {
                    best_match = &male_manatees[j];
				}
				// Once we find one that's too big or too old, break off.
				if(male_manatees[j].age > youngest || best_match.size > first_female.size) {
					break;
				}
			}

			if(best_match.size > first_female.size) {
				// There's no good match for this female.
				// The setup is impossible.
				impossible = true;
				break;
			}
			// Otherwise, add the pair.
			female_answer.push(female_manatees[i].index);
			male_answer.push(best_match.index);
		}
	} 
	//if impossible {
	//	print!("impossible");	
	//}
	
	
	//let mut manatee_index = male_manatees.binary_search_by_key(&11, |&Manatee{ size, age, index }| age).unwrap();
	//print!("{}", male_manatees[manatee_index].size);
	
	// print indices of manatees
	print!("\n");
	for i in 0..(pairs) {
		let current_manatee = female_answer[i];
		print!("{} ", current_manatee);
	}
	print!("\n");
	for i in 0..(pairs) {
		let current_manatee = male_answer[i];
		print!("{} ", current_manatee);
	}
	
	
	
	// Print the manatee sets
	print!("\n");
	for i in 0..(pairs) {
		let current_manatee = female_manatees[i];
		print!("{} ", current_manatee.age);
	}
	print!("\n");
	for i in 0..(pairs) {
		let current_manatee = female_manatees[i];
		print!("{} ", current_manatee.size);
	}
	
	
	print!("\n\n");
	for i in 0..(pairs) {
		let current_manatee = male_manatees[i];
		print!("{} ", current_manatee.age);
	}
	print!("\n");
	for i in 0..(pairs) {
		let current_manatee = male_manatees[i];
		print!("{} ", current_manatee.size);
	}
	Ok(())
}