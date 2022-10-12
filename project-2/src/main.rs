use std::io;

pub struct Manatee {
	size: i32,
	age: i32,
	index: i32
}

fn main() -> io::Result<()> {
	let mut pairs: i32;
	let mut female_manatees: Vec<Manatee>;
	let mut male_manatees: Vec<Manatee>;
	
	// Read the first five lines from stdio
	for n in 1..6  {
		// Set the input variable to be a new string
		let mut buffer = String::new();
		// Set the next input equal to the input variable
		io::stdin().read_line(&mut buffer)?;
		
		
		match n {
			// Get number of pairs
			1 => pairs = n,
			
			// Get female manatee age
			2 => 
			for i in 1..pairs {
				female_manatees[i].age = buffer[i].parse::<i32>().unwrap();
			},
			// Get female manatee size
			3 =>
			for i in 1..pairs {
				female_manatees[n].size = buffer[i].parse::<i32>().unwrap();
			},
			// Get male manatee age
			4 =>
			for i in 1..pairs {
				male_manatees[n].age = buffer[i].parse::<i32>().unwrap();
			},
			// Get male manatee size
			5 => 
			for i in 1..pairs {
				male_manatees[n].size = buffer[i].parse::<i32>().unwrap();
			},
		}			
	}
	Ok(())
}