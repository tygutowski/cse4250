package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"log"
	"strconv"
)

type Manatee struct {
	age   int
	size  int
	index int
}

func (m *Manatee) SetAge(newAge int) {
	m.age = newAge
}

func (m Manatee) GetAge() int {
	return (m.age)
}

func (m *Manatee) SetSize(newSize int) {
	m.size = newSize
}

func (m Manatee) GetSize() int {
	return (m.size)
}

func (m *Manatee) SetIndex(newIndex int) {
	m.index = newIndex
}

func (m Manatee) GetIndex() int {
	return (m.index)
}

func SortByAge(m []Manatee) []Manatee {
	sort.Slice(m, func(i, j int) bool { return m[i].GetAge() < m[j].GetAge() })
	return m
}

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
	
	fmt.Printf("There are %d pairs of manatees\n", pairs)
	
	
	// Instance an array for both male and female manatees
	fManatees := make([]Manatee, pairs)
	mManatees := make([]Manatee, pairs)
	for i := 0; i < 4; i++ {
		fManatee := new(Manatee)
		mManatee := new(Manatee)
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
		mManatees[j].SetSize(mSize)
	}
	
	
	for j := 0; j < pairs; j++ {
		fmt.Printf("Male manatee %d is %d years old, and is a size of %d\n", j, mManatees[j].GetAge(), mManatees[j].GetSize())
	}
	// Sort arrays
	//SortByAge(fManatees)
	//SortByAge(mManatees)
	
	// Print the age of all of the female manatees
	//for i := range fManatees {
	//	fmt.Println(fManatees[i].GetAge())
	//}

	// Print the age of all of the male manatees
	//for i := range mManatees {
	//	fmt.Println(mManatees[i].GetAge())
	//}

}