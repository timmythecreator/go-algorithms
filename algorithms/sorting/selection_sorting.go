package sorting

import "fmt"

func SelectionSortSample() {
	fmt.Println(`
	Selection sorting:
The goal is to sort by finding the smallest element and swapping it with the first unsorted element.
This process repeats for the remaining unsorted elements until the array is sorted.
	`)
	arr := []int{5, 3, 8, 4, 2}
	fmt.Println("\nUnsorted array:", arr)
	fmt.Print(`
For example, let's observe the first pass of sorting:
5>3 -> minIndex = 1, 3<8 -> minIndex = 1,
3<4 -> minIndex = 1, 3>2 -> minIndex = 4
minIndex != i -> 5 and 2 are swapped
So this is how it will look after the first pass of sorting:
[2, 3, 8, 4, 5]
	`)
	SelectionSort(arr)
	fmt.Println("\nSorted array:", arr)
	fmt.Println("\nThe complexity of this sorting algorithm is O(n^2)")
}

func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
