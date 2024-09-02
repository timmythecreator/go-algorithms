package searching

import "fmt"

func TwoSumSample() {
	fmt.Print(`
	TwoSum:
The goal is to find two numbers in an array that add up to a given target sum.
We changed the initial version of the TwoSum function to a more optimized version.
The new version uses two pointers to iterate over the array and find the target sum.
Why is this version better?
Because it doesn't use a hash map to store the numbers we have seen so far.
	`)
	arr := []int{2, 3, 4, 5, 7}
	target := 9
	fmt.Println("\nSample array (sorted):", arr)
	fmt.Println("Target sum:", target)
	result := TwoSum(arr, target)
	fmt.Print(`
So, how will it look like:
We iterate over the array and for each number, we calculate its complement.
If the complement exists in the hash map, we return the indices of the current number and its complement.
And this is how it will look like:`)
	fmt.Print(result)
	fmt.Println("\nThe runtime of this algorithm is O(n), the space complexity is O(1).")
}

func TwoSum(arr []int, target int) []int {
	i, r := 0, len(arr)-1
	for i < r {
		sum := arr[i] + arr[r]
		if sum == target {
			return []int{i, r}
		}
		if sum < target {
			i++
		} else {
			r--
		}
	}
	return nil
}

// This is the first version of the TwoSum function.
// It uses a hash map to store the numbers we have seen so far and their indices.
/*
func TwoSum(arr []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range arr {
		complement := target - num
		if j, exists := numMap[complement]; exists {
			return []int{i, j}
		}
		numMap[num] = i
	}
	return nil
}
*/
