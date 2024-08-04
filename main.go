package main

import (
	"fmt"
	"go-algorithms/algorithms"
	"go-algorithms/algorithms/array"
	datastructures "go-algorithms/algorithms/data_structures"
	"go-algorithms/algorithms/recursive"
	"go-algorithms/algorithms/searching"
	"go-algorithms/algorithms/sorting"
)

var algs = []algorithms.Algorithms{
	{Name: "Bubble Sort", Run: sorting.BubbleSortSample},
	{Name: "Selection Sort", Run: sorting.SelectionSortSample},
	{Name: "Quick Sort", Run: sorting.QuickSortSample},
	{Name: "Minimum Value", Run: searching.MinimumValueSample},
	{Name: "Maximum Value", Run: searching.MaximumValueSample},
	{Name: "Binary Search", Run: searching.BinarySearchSample},
	{Name: "Two Sum", Run: searching.TwoSumSample},
	{Name: "Reverse Array", Run: array.ReverseArraySample},
	{Name: "Factorial", Run: recursive.FactorialSample},
	{Name: "Fibonacci", Run: recursive.FibonacciSample},
	{Name: "Queue", Run: datastructures.QueueSample},
	{Name: "Stack", Run: datastructures.StackSample},
}

func main() {
	fmt.Println("Available Algorithms:")
	for i, algorithm := range algs {
		fmt.Printf("%d: %s\n", i, algorithm.Name)
	}

	var choice int
	fmt.Print("Enter the number of the algorithm to run: ")
	fmt.Scan(&choice)

	if choice >= 0 && choice < len(algs) {
		algs[choice].Run()
	} else {
		fmt.Println("Invalid choice")
	}
}
