package algorithms

// Algorithm represents an algorithm with its name and function
type Algorithms struct {
	Name string
	Run  func()
}
