/*
 * Author:  Tyler Gutowski, Remington Greko: tgutowski2020@my.fit.edu, rgreko2020@my.fit.edu
 * Course:  CSE 4250, Section 01, Fall 2022
 * Project: Proj 01, Manatee Evacuation
 */

package main

// Import everything needed
import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var mPermutations [][]Manatee
var fPermutations [][]Manatee

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

// Checks the order of Manatees to see if they're in a good order.
// Returns "impossible" if it isn't currently possible, otherwise
// returns the order of manatees as a string.
func CheckOrder(f []Manatee, m []Manatee, pairs int) string {
	// Check to see if front is smaller than back.
	for i := 0; i < pairs; i++ {
		// If at any point it isn't, break
		if f[i].GetSize() >= m[i].GetSize() {
			return ""
		}
	}
	// Otherwise, it's good!
	var answer = ""
	for i := 0; i < pairs; i++ {
		answer = answer + strconv.Itoa(f[i].GetIndex()) + " "
	}
	answer = answer + "\n"
	for i := 0; i < pairs; i++ {
		answer = answer + strconv.Itoa(m[i].GetIndex()) + " "
	}
	return answer
}

// Checks if the number is valid (Between 1 and 10^9, inclusive)
// Returns an error, or nil if there is no error.
func IsNumberValid(num int) error {
	if num < 1 || num > 1000000000 {
		return errors.New("Number invalid. Bounds are {1 <= n <= 10^9}")
	}
	return nil
}

func sliceBetween(begin, end int) []int {
	if begin > end {
		panic("Invalid slice range!")
	}
	slice := make([]int, end-begin)
	for i := 0; i < len(slice); i++ {
		slice[i] = i + 1 + begin
	}
	return slice
}

// Perm calls f with each permutation of a.
//func Perm(a []Manatee, f func([]Manatee)) {
//	perm(a, f, 0)
//}

// Permute the values at index i to len(a)-1.
func mHeapPermutation(a []Manatee, size int, n int, fPerms [][]Manatee) {
	if size == 1 {
		a2 := make([]Manatee, len(a))
		copy(a2, a)
		mPermutations = append(mPermutations, a2)
	}
	for i := 0; i < size; i++ {
		mHeapPermutation(a, size-1, n, fPerms)
		if size%2 == 1 {
			temp := a[0]
			a[0] = a[size-1]
			a[size-1] = temp
		} else {
			temp := a[i]
			a[i] = a[size-1]
			a[size-1] = temp
		}
	}
}

func fHeapPermutation(a []Manatee, size int, n int, fPerms [][]Manatee) {
	if size == 1 {
		a2 := make([]Manatee, len(a))
		copy(a2, a)
		fPermutations = append(fPermutations, a2)
	}
	for i := 0; i < size; i++ {
		fHeapPermutation(a, size-1, n, fPerms)
		if size%2 == 1 {
			temp := a[0]
			a[0] = a[size-1]
			a[size-1] = temp
		} else {
			temp := a[i]
			a[i] = a[size-1]
			a[size-1] = temp
		}
	}
}

// Main runner function.
func main() {
	// Check the parameters. We need 2, one for execution, one for input file.
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	// Try to open the file.
	f, err := os.Open(os.Args[1])
	// If there is an error, log it.
	if err != nil {
		log.Fatalln(err)
	}
	// Close the file upon ending main()
	defer f.Close()

	// Instance a scanner for inputs
	scanner := bufio.NewScanner(f)

	// Scan the file
	scanner.Scan()

	// Take the first value, and make it represent each pair of manatees.
	pairs, err := strconv.Atoi(scanner.Text())
	// Check if there's an error. This will occur if the value isn't an integer.
	if err != nil {
		log.Fatalln(err)
	}
	err2 := IsNumberValid(pairs)
	if err2 != nil {
		log.Fatalln(err2)
	}

	// Instance an array for both male and female manatees
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

	// While the scanner can still scan
	scanner.Scan()
	line := scanner.Text()
	numbers := strings.Fields(line)
	for j := 0; j < pairs; j++ {
		fAge, err := strconv.Atoi(numbers[j])
		if err != nil {
			log.Fatalln(err)
		}
		err2 := IsNumberValid(fAge)
		if err2 != nil {
			log.Fatalln(err2)
		}
		fManatees[j].SetAge(fAge)
	}
	scanner.Scan()
	line = scanner.Text()
	numbers = strings.Fields(line)
	for j := 0; j < pairs; j++ {
		fSize, err := strconv.Atoi(numbers[j])
		if err != nil {
			log.Fatalln(err)
		}
		err2 := IsNumberValid(fSize)
		if err2 != nil {
			log.Fatalln(err2)
		}
		fManatees[j].SetSize(fSize)
	}
	scanner.Scan()
	line = scanner.Text()
	numbers = strings.Fields(line)
	for j := 0; j < pairs; j++ {
		mAge, err := strconv.Atoi(numbers[j])
		if err != nil {
			log.Fatalln(err)
		}
		err2 := IsNumberValid(mAge)
		if err2 != nil {
			log.Fatalln(err2)
		}
		mManatees[j].SetAge(mAge)
	}
	scanner.Scan()
	line = scanner.Text()
	numbers = strings.Fields(line)
	for j := 0; j < pairs; j++ {
		mSize, err := strconv.Atoi(numbers[j])
		if err != nil {
			log.Fatalln(err)
		}
		err2 := IsNumberValid(mSize)
		if err2 != nil {
			log.Fatalln(err2)
		}
		mManatees[j].SetSize(mSize)
	}

	SortByAge(fManatees)
	SortByAge(mManatees)
	// PERMUTATE THROUGH EACH POSSIBLE COMBINATION
	///for {

	fSizes := make([]int, pairs)
	mSizes := make([]int, pairs)
	for i := 0; i < pairs; i++ {
		fSizes[i] = fManatees[i].GetSize()
		mSizes[i] = mManatees[i].GetSize()
	}
	//found := false

	emptySlice := [][]Manatee{}
	fHeapPermutation(fManatees, pairs, pairs, emptySlice)
	mHeapPermutation(mManatees, pairs, pairs, emptySlice)
	
	fmt.Println(fPermutations)
	fmt.Println(mPermutations)
	
	for i := 0; i < 
}