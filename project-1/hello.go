package main

import (
	//"io/ioutil"
	"fmt"
	"os"
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
	return(m.age)
}

func (m Manatee) SetSize(size int) {
	m.size = size
}

func (m Manatee) GetSize() int {
	return(m.size)
}

func (m Manatee) SetIndex(index int) {
	m.index = index
}

func (m Manatee) GetIndex() int {
	return(m.index)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
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
