package sorting

import "fmt"

// QuickSortSample demonstrates quick sort algorithm
func QuickSortSample() {
	fmt.Print(`
	Quick sorting:
The goal is to sort an array by dividing it into partitions around a pivot element.
Elements less than the pivot go to the left, elements greater than the pivot go to the right.")
This process continues recursively for each partition.
	`)
	fmt.Println()

	arr := []int{5, 3, 8, 4, 2}
	fmt.Println("Unsorted array:", arr)
	fmt.Print(`
So, how will it look like:
1. Choose a pivot element (e.g., the last element).
2. Partition the array such that elements less than the pivot are on the left, and elements greater than the pivot are on the right.
3. Recursively apply the same process to the left and right subarrays.
And this is how it will look like:
	`)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
	fmt.Println("\nThe complexity of this algorithm is O(n log n) on average.")
}

// QuickSort sorts an array using quick sort algorithm
func QuickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

// partition partitions the array and returns the index of the pivot
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
