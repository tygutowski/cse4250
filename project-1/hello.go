/*
 * Author:  Remington Greko, Tyler Gutowski: rgreko2020@my.fit.edu, tgutowski2020@my.fit.edu
 * Course:  CSE 4250, Section 01, Fall 2022
 * Project: Proj 01, Manatee Evacuation 
 */

package main

import (
	"complex"
	"errors"
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

func CheckOrder(f []Manatee, m []Manatee, pairs int) string {
	// Check to see if front is smaller than back.
	for i := 0; i < pairs; i++ {
		// If at any point it isn't, break
		if f[i].GetSize() >= m[i].GetSize() {
			
			return "impossible"
		}
	}
	// Otherwise, it's good!
	fString := strings.Join(f.GetAge," ")
	mString := strings.Join(m.GetSize," ")
	answer := fString + "\n" + mString
	return answer
}

func IsNumberValid(num int) (int, error) {
	if num < 1 || num > 1000000000 {
		return 0, errors.New("Number invalid. Bounds are {1 <= n <= 10^9}")
	}
	return num, nil
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
	IsNumberValid(pairs)
		
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

	SortByAge(fManatees)
	SortByAge(mManatees)
	
	answer := CheckOrder(fManatees, mManatees, pairs)
	if answer == "impossible" {
		fmt.Println("impossible")
	} else {
		fmt.Println(answer)
	}
	
	// Print the age of all of the female manatees
	//for i := range fManatees {
	//	fmt.Println(fManatees[i].GetAge())
	//}

	// Print the age of all of the male manatees
	//for i := range mManatees {
	//	fmt.Println(mManatees[i].GetAge())
	//}

}