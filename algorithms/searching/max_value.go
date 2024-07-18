package searching

import "fmt"

func MaximumValueSample() {
	fmt.Print(`
	Maximum value:
The goal is to find the maximum value in an array by comparing each element.
This process repeats for all elements in the array to find the biggest value.
	`)
	arr := []int{5, 3, 8, 4, 2}
	fmt.Println("\nSample array:", arr)
	fmt.Print(`
For example, let's observe the pass:
Setup maxIdx = 0
Then compare each element with the maxIdx element:
5>3 -> maxIdx = 0, 5<8 -> maxIdx = 2,
8>4 -> maxIdx = 2, 8>2 -> maxIdx = 2
`)
	maxValue := MaximumValue(arr)
	fmt.Printf("return value of maxIdx which is %d\n", maxValue)
	fmt.Println("\nThe complexity of this searching algorithm is O(n)")
}

func MaximumValue(arr []int) int {
	n := len(arr)
	if n == 0 {
		fmt.Println("Array is empty")
		return -1
	}
	maxIdx := 0
	for i := 1; i < n; i++ {
		if arr[maxIdx] < arr[i] {
			maxIdx = i
		}
	}
	return arr[maxIdx]
}
