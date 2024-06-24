package searching

import "fmt"

func TwoSumSample() {
	fmt.Print(`
	TwoSum:
The goal is to find two numbers in an array that add up to a given target sum.
We use a hash map to store the numbers we have seen so far and their indices.
For each number, we calculate its complement with respect to the target sum.
If the complement exists in the hash map, we return the indices of the current number and its complement.
	`)
	arr := []int{2, 3, 7, 4, 5}
	target := 9
	fmt.Println("\nSample array:", arr)
	fmt.Println("Target sum:", target)
	result := TwoSum(arr, target)
	fmt.Print(`
So, how will it look like:
We iterate over the array and for each number, we calculate its complement.
If the complement exists in the hash map, we return the indices of the current number and its complement.
And this is how it will look like:`)
	fmt.Print(result)
	fmt.Println("\nThe complexity of this algorithm is O(n)")
}

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
