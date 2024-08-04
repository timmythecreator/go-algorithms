package datastructures

import "fmt"

func StackSample() {
	fmt.Print(`
	Stack:
Stack is a linear data structure that follows the LIFO (Last In First Out) principle.
The elements are inserted and deleted from the same end.
Let's observe the example below:`)
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println("\nStack after pushing 1, 2, 3:", s.items)

	popped := s.Pop()
	fmt.Println("Popped item:", popped)
	fmt.Println("Stack after popping:", s.items)
	fmt.Println("The complexity of this data structure is O(1) for both push and pop operations.")
}

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}
