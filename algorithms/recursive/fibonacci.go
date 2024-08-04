package recursive

import "fmt"

func FibonacciSample() {
	fmt.Print(`
	Fibonacci:
Fibonacci is a sequence of numbers in which each number is the sum of the two preceding ones.
Let's observe the example below:`)
	n := 10
	fib := Fibonacci(n)
	fmt.Println("\nThe", n, "th Fibonacci number is", fib)
	fmt.Println("The complexity of this recursive algorithm is O(2^n).")
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
