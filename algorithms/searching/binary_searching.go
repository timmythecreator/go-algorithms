package searching

import "fmt"

func BinarySearchSample() {
	fmt.Print(`
	Binary Search:
The goal is to find an element in a sorted array using binary search.
This process divides the search interval in half repeatedly.
	`)
	arr := []int{2, 3, 4, 5, 8}
	target := 4
	fmt.Println("\nSorted array:", arr)
	fmt.Printf("Searching for %d...\n", target)
	index := BinarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
	fmt.Println("\nThe complexity of this algorithm is O(log n)")
}

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
