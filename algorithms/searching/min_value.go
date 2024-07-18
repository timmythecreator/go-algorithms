package searching

import "fmt"

func MinimumValueSample() {
	fmt.Print(`
	Minimum value:
The goal is to find the minimum value in an array by comparing each element.
This process repeats for all elements in the array to find the smallest value.
	`)
	arr := []int{5, 3, 8, 4, 2}
	fmt.Println("\nSample array:", arr)
	fmt.Print(`
For example, let's observe the pass:
Setup minIdx = 0
Then compare each element with the minIdx element:
5>3 -> minIdx = 1, 3<8 -> minIdx = 1,
3<4 -> minIdx = 1, 2<3 -> minIdx = 4
`)
	minValue := MinimumValue(arr)
	fmt.Printf("return value of minIdx which is %d\n", minValue)
	fmt.Println("\nThe complexity of this searching algorithm is O(n)")
}

func MinimumValue(arr []int) int {
	n := len(arr)
	if n == 0 {
		fmt.Println("Array is empty")
		return -1
	}
	minIdx := 0
	for i := 1; i < n; i++ {
		if arr[minIdx] > arr[i] {
			minIdx = i
		}
	}
	return arr[minIdx]
}
