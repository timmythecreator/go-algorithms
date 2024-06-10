package main

import (
	"algorithms/sorting"
	"fmt"
)

// Algorithm represents an algorithm with its name and function
type Algorithm struct {
	Name string
	Run  func()
}

var algorithms = []Algorithm{
	{"Bubble Sort", sorting.BubbleSortExample},
}

func main() {
	fmt.Println("Available Algorithms:")
	for i, algorithm := range algorithms {
		fmt.Printf("%d: %s\n", i, algorithm.Name)
	}

	var choice int
	fmt.Print("Enter the number of the algorithm to run: ")
	fmt.Scan(&choice)

	if choice >= 0 && choice < len(algorithms) {
		algorithms[choice].Run()
	} else {
		fmt.Println("Invalid choice")
	}
}
