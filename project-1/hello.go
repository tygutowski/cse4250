package main

import (
	//"io/ioutil"
	"bufio"
	"fmt"
	"os"
	"sort"
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

func ReadInput(fpath string) ([]string, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	/// Idk how the fuck to do input please help
	ReadInput(os.Args[0])

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
	//data, err := ioutil.ReadFile(os.Args[1])
	//if err != nil {
	//	fmt.Println("Cannot read file:", os.Args[1])
	//	panic(err)
	//}
	//for _, line := range data.slice {
	//	fmt.Println(line)
	//}

	// Example manatee instance
	m := Manatee{age: 4, size: 6, index: 0}
	fmt.Println(m.age)
}
