use std::io;

pub struct Manatee {
	size: i32,
	age: i32,
	index: i32
}

fn main() -> io::Result<()> {
	let mut pairs: usize = 0;
	let mut female_manatees: Vec<Manatee> = Vec::new();
	let mut male_manatees: Vec<Manatee> = Vec::new();
	
	// Read the first five lines from stdio
	for n in 1..6  {
		// Set the input variable to be a new string
		let mut buffer = String::new();
		// Set the next input equal to the input variable
		io::stdin().read_line(&mut buffer)?;
		let mut buffer_split = buffer.split_whitespace();
		
		match n {
			// Get number of pairs
			1 => pairs = buffer_split.next().expect("found empty string where number expected").parse::<usize>().unwrap(),
			
			// Get female manatee age
			2 => 
			for i in 1..pairs {
				let ind: i32 = i as i32;
				let empty_manatee = Manatee {
					size: 0,
					age: 0,
					index: ind
				};
				female_manatees.push(empty_manatee);
				female_manatees[i].age = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
			},
			// Get female manatee size
			3 =>
			for i in 1..pairs {
				female_manatees[i].size = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
			},
			// Get male manatee age
			4 =>
			for i in 1..pairs {
				let ind: i32 = i as i32;
				let empty_manatee = Manatee {
					size: 0,
					age: 0,
					index: ind
				};
				male_manatees.push(empty_manatee);
				male_manatees[i].age = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
			},
			// Get male manatee size
			5 => 
			for i in 1..pairs {
				male_manatees[i].size = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
			},
		}			
	}

	// Now that we have the trees, sort by age
    // First the female manatees...
    for j in 1..(pairs - 1) {
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
			if male_manatees[n].age <= male_manatees[swap_index].age {
				swap_index = n;
			}
		}
        let temp = male_manatees[j];
		male_manatees[j] = male_manatees[swap_index];
		male_manatees[swap_index] = temp;
	}

	Ok(())
}