/*
 * Author:  Remington Greko, Tyler Gutowski: rgreko2020@my.fit.edu, tgutowski2020@my.fit.edu
 * Course:  CSE 4250, Section 01, Fall 2022
 * Project: Proj 01, Manatee Evacuation 
 */

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

func (m Manatee) SetAge(age int) {
	m.age = age
}

func (m Manatee) GetAge() int {
	return (m.age)
}

func (m Manatee) SetSize(size int) {
	m.size = size
}

func (m Manatee) GetSize() int {
	return (m.size)
}

func (m Manatee) SetIndex(index int) {
	m.index = index
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
		fManatees[j].age, err = strconv.Atoi(numbers[j])
	}
	scanner.Scan()
	line = scanner.Text()
	numbers = strings.Fields(line)
	for j := 0; j < pairs; j++ {
		fManatees[j].size, err = strconv.Atoi(numbers[j])
	}
	scanner.Scan()
	line = scanner.Text()
	numbers = strings.Fields(line)
	for j := 0; j < pairs; j++ {
		mManatees[j].age, err = strconv.Atoi(numbers[j])
	}
	scanner.Scan()
	line = scanner.Text()
	numbers = strings.Fields(line)
	for j := 0; j < pairs; j++ {
		mManatees[j].size, err = strconv.Atoi(numbers[j])
	}
	
	
	for j := 0; j < pairs; j++ {
		fmt.Printf("Male manatee %d is %d years old, and is a size of %d\n", j, mManatees[j].age, mManatees[j].size)
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