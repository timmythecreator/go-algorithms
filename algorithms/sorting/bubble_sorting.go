package sorting

import "fmt"

func BubbleSortSample() {
	fmt.Print(`
	Bubble sorting:
The goal is to sort by comparing arrays[j] with arrays[j+1],
and if arrays[j] > arrays[j+1], they're swapping.
The remaining elements repeat this process until sorting is completed.
	`)
	arr := []int{5, 3, 8, 4, 2}
	fmt.Println("\nUnsorted array:", arr)
	fmt.Print(`
For example, let's observe the first pass of sorting:
5>3 -> swap, 5<8 -> no swap,
8>4 -> swap, 8>2 -> swap
So this is how it will look after the first pass of sorting:
[3, 5, 4, 2, 8]
	`)
	BubbleSort(arr)
	fmt.Println("\nSorted array:", arr)
	fmt.Println("\nThe complexity of this sorting algorithm is O(n^2)")
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
