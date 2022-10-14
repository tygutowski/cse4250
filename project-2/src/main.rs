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
	let mut max_female_size: i32 = 0;
	let mut max_male_size: i32 = 0;
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

	
	let mut impossible: bool = true;
	
	let mut female_answer: Vec<i32> = Vec::new();
	let mut male_answer: Vec<i32> = Vec::new();
	
	// morbin' time.

	for i in 0..pairs-1 {
		// Sort both lists of manatees
		female_manatees.sort_by(|a, b| a.cmp(b));
		male_manatees.sort_by(|a, b| a.cmp(b));
		// Get the first manatee in each list
		let mut first_male = &male_manatees[0];
		let mut first_female = &female_manatees[0];
		// Get the smallest (first) female that is larger than the male.
		if(female_manatees[i].size > first_male.size) {
			print!("{}:{}", female_manatees[i].age, female_manatees[i].size);
			female_answer.push(female_manatees[i].index);
			male_answer.push(male_manatees[i].index);
			break;
		}
	}
	//if impossible {
	//	print!("impossible");	
	//}
	
	
	//let mut manatee_index = male_manatees.binary_search_by_key(&11, |&Manatee{ size, age, index }| age).unwrap();
	//print!("{}", male_manatees[manatee_index].size);
	
	// print indeces of manatees
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