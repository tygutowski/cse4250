package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	//"io/ioutil"
	"log"
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

func ReadFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	ReadFile(os.Args[1])

	var fem_manatees []Manatee
	var male_manatees []Manatee

	SortByAge(fem_manatees)
	SortByAge(male_manatees)
	for i := range fem_manatees {
		fmt.Println(fem_manatees[i].GetAge())
	}
	print("\n")
	for i := range male_manatees {
		fmt.Println(male_manatees[i].GetAge())
	}

}