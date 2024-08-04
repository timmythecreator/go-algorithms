package recursive

import "fmt"

func FactorialSample() {
	fmt.Print(`
	Factorial:
Factorial of a number is the product of all positive integers less than or equal to n.
Let's observe the example below:
`)
	n := 5
	fmt.Println("The factorial of", n, "is ", Factorial(n))
	fmt.Println("(5*4*3*2*1).\nThe complexity of this recursive algorithm is O(n).")
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
