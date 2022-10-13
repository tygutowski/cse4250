use std::io;
use std::collections::BTreeSet;
use std::cmp::Ordering;

// Manatee struct
#[derive(Debug)]
pub struct Manatee {
	size: i32,
	age: i32,
	index: i32
}
impl Ord for Manatee {
	fn cmp(&self, other: &Self) -> Ordering {
		if self.get_age() == other.get_age() {
			if self.get_size() == other.get_size() {
				return self.get_index().cmp(&other.get_index());
			}
			return self.get_size().cmp(&other.get_size());
		}
		return self.get_age().cmp(&other.get_age());
	}
}
impl Manatee {
	fn get_size(&self) -> i32 {
		return self.size;
	}
	fn get_age(&self) -> i32 {
		return self.age;
	}
	fn get_index(&self) -> i32 {
		return self.index;
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
				let empty_manatee = Manatee {
					size: 0,
					age: 0,
					index: ind
				};
				female_manatees.insert(empty_manatee);
				female_manatees[i-1].age = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
			},
			
			// Get female manatee size
			3 =>
			for i in 1..(pairs+1) {
				female_manatees[i-1].size = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
			},
			
			// Get male manatee age
			4 =>
			for i in 1..(pairs+1) {
				let ind: i32 = i as i32;
				let empty_manatee = Manatee {
					size: 0,
					age: 0,
					index: ind
				};
				male_manatees.insert(empty_manatee);
				male_manatees[i-1].age = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
			},
			
			// Get male manatee size
			5 => 
			for i in 1..(pairs+1) {
				male_manatees[i-1].size = buffer_split.next().expect("found empty string where number expected").parse::<i32>().unwrap();
			},
			
			// For anything else (which shouldn't occur)
			_ =>
			break,
		}			
	}

	// Now that we have the trees, sort by age
	female_manatees.sort_by_key(|a| a.age);
	print!("Female Manatees: \n");
	for i in 1..(pairs+1) {
		print!("S: {}  A: {}  I: {}\n", female_manatees[i-1].get_size(), female_manatees[i-1].get_age(), female_manatees[i-1].get_index());
	}
	print!(" ");
	male_manatees.sort_by_key(|a| a.age);
	print!("Male Manatees:   \n");
	for i in 1..(pairs+1) {
		print!("S: {}  A: {}  I: {}\n", male_manatees[i-1].get_size(), male_manatees[i-1].get_age(), male_manatees[i-1].get_index());
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