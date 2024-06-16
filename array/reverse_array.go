package array

import "fmt"

func ReverseArraySample() {
	fmt.Print(`
	ReverseArray:
The goal is to reverse an array by swapping the first element with the last and repeating this process until reaching the middle.
This process continues until all elements are swapped.
	`)
	arr := []int{2, 3, 4, 5, 8}
	fmt.Println("\nSample array:", arr)
	ReverseArray(arr)
	fmt.Print(`
Let's observe the algorithm more detailed:
We initialize a temporary value for our last element,
then we save the first element in the last position,
after that we take the value from the temporary variable and place it in the first position.
And this is how it will look like after all: `)
	fmt.Print(arr)
	fmt.Println("\nThe complexity of this algorithm is O(n)")
}

func ReverseArray(arr []int) {
	n := len(arr)
	mid := n / 2
	for i := 0; i <= mid; i++ {
		var temp = arr[i]
		arr[i] = arr[n-i-1]
		arr[n-i-1] = temp
	}
}
