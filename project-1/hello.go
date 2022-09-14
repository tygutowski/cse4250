package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Cannot read file:", os.Args[1])
		panic(err)
	}
	for _, line := range data.slice {
		fmt.Println(line)
	}
}
