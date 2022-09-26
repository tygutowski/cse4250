/*
 * Author:  Tyler Gutowski, Remington Greko: tgutowski2020@my.fit.edu, rgreko2020@my.fit.edu
 * Course:  CSE 4250, Section 01, Fall 2022
 * Project: Proj 01, Manatee Evacuation
 */

package main

// Import all necessary packages
import (
	"fmt"
	"log"
	"sort"
)

var mPermutations [][]Manatee /// Slice of male manatee permutations
var fPermutations [][]Manatee /// Slice of female manatee permutations

// Create Manatee struct with age, size and initial index
type Manatee struct {
	age   int
	size  int
	index int
}

// Setter for age
func (m *Manatee) SetAge(newAge int) {
	m.age = newAge
}

// Getter for age
func (m Manatee) GetAge() int {
	return (m.age)
}

// Setter for size
func (m *Manatee) SetSize(newSize int) {
	m.size = newSize
}

// Getter for size
func (m Manatee) GetSize() int {
	return (m.size)
}

// Setter for index
func (m *Manatee) SetIndex(newIndex int) {
	m.index = newIndex + 1
}

// Getter for index
func (m Manatee) GetIndex() int {
	return (m.index)
}

// Sorting function to organize a slice of Manatees by their ages.
func SortByAge(m []Manatee) []Manatee {
	sort.Slice(m, func(i, j int) bool { return m[i].GetAge() < m[j].GetAge() })
	return m
}

// Permute the male list, mPermutations.
func mHeapPermutation(validPerm []Manatee, size int, n int, fPerms [][]Manatee) {
	if size == 1 {
		a2 := make([]Manatee, len(validPerm))
		copy(a2, validPerm)
		mPermutations = append(mPermutations, a2)
	}
	for i := 0; i < size; i++ {
		mHeapPermutation(validPerm, size-1, n, fPerms)
		if size%2 == 1 {
			temp := validPerm[0]
			validPerm[0] = validPerm[size-1]
			validPerm[size-1] = temp
		} else {
			temp := validPerm[i]
			validPerm[i] = validPerm[size-1]
			validPerm[size-1] = temp
		}
	}
}

// Permute the female list, fPermutations.
func fHeapPermutation(validPerm []Manatee, size int, n int, fPerms [][]Manatee) {
	if size == 1 {
		a2 := make([]Manatee, len(validPerm))
		copy(a2, validPerm)
		fPermutations = append(fPermutations, a2)
	}
	for i := 0; i < size; i++ {
		fHeapPermutation(validPerm, size-1, n, fPerms)
		if size%2 == 1 {
			temp := validPerm[0]
			validPerm[0] = validPerm[size-1]
			validPerm[size-1] = temp
		} else {
			temp := validPerm[i]
			validPerm[i] = validPerm[size-1]
			validPerm[size-1] = temp
		}
	}
}

// Main runner function.
func main() {
	// Get the number of pairs from std input
	var pairs int
	_, err := fmt.Scanf("%d", &pairs)
	if err != nil {
		log.Fatalln(err)
	}
	
	/// Instance an array for both male and female manatees
	fManatees := make([]Manatee, pairs)
	mManatees := make([]Manatee, pairs)
	for i := 0; i < pairs; i++ {
		fManatee := new(Manatee)
		fManatee.SetIndex(i)
		mManatee := new(Manatee)
		mManatee.SetIndex(i)
		fManatees[i] = *fManatee
		mManatees[i] = *mManatee
	}

	/// Scan everything from the input and put it into its respective
	/// categories. fAge, fSize, mAge and mSize.
	/// Scan for j amount of integers, depending on how many pairs there are.
	for j := 0; j < pairs; j++ {
		var fAge int
		_, err := fmt.Scanf("%d", &fAge)
		if err != nil {
			log.Fatalln(err)
		}
		fManatees[j].SetAge(fAge)
	}
	for j := 0; j < pairs; j++ {
		var fSize int
		_, err := fmt.Scanf("%d", &fSize)
		if err != nil {
			log.Fatalln(err)
		}
		fManatees[j].SetSize(fSize)
	}
	for j := 0; j < pairs; j++ {
		var mAge int
		_, err := fmt.Scanf("%d", &mAge)
		if err != nil {
			log.Fatalln(err)
		}
		mManatees[j].SetAge(mAge)
	}

	for j := 0; j < pairs; j++ {
		var mSize int
		_, err := fmt.Scanf("%d", &mSize)
		if err != nil {
			log.Fatalln(err)
		}
		mManatees[j].SetSize(mSize)
	}

	/// Instance an empty slice, then run to generate all permutations for
	/// male and female manatees.
	emptySlice := [][]Manatee{}
	fHeapPermutation(fManatees, pairs, pairs, emptySlice)
	mHeapPermutation(mManatees, pairs, pairs, emptySlice)

	/// Variable to see if the input will work
	willWork := true
	/// Set break point
	looper:
	/// For all female permutations
	for fPermIndex := 0; fPermIndex < len(fPermutations); fPermIndex++ {
		/// For all male permutations
		for mPermIndex := 0; mPermIndex < len(mPermutations); mPermIndex++ {
			willWork = true
			/// Check the individual manatees
			for structIndex := 0; structIndex < pairs; structIndex++ {
				/// If the size of the female is larger than the size of the male
				/// then the code executes properly.
				if (fPermutations[fPermIndex][structIndex].GetSize() > mPermutations[mPermIndex][structIndex].GetSize()) {
					
				} else {
					willWork = false
					break
				}
			}
			/// If the age of each concurrent manatee is equal or increasing,
			/// then the code executes property.
			for structIndex := 0; structIndex < pairs - 1; structIndex++ {
				if fPermutations[fPermIndex][structIndex].GetAge() <= fPermutations[fPermIndex][structIndex + 1].GetAge() {
					
				} else {
					willWork = false
					break
				}
				if mPermutations[mPermIndex][structIndex].GetAge() <= mPermutations[mPermIndex][structIndex + 1].GetAge() {
					
				} else {
					willWork = false
					break
				}
			}
			/// If all previous tests work, you pass! YAY!
			/// Print female and male indeces, then break to break point.
			if willWork {
				for k := 0; k < pairs; k++ {
					fmt.Printf("%d ",fPermutations[fPermIndex][k].GetIndex())
				}
				fmt.Println()
				for k := 0; k < pairs; k++ {
					fmt.Printf("%d ",mPermutations[mPermIndex][k].GetIndex())
				}
				fmt.Println()
				break looper
			}
		}
	}
	/// If it never ended up working, then its impossible.
	if !willWork {
		print("impossible")
	}
}