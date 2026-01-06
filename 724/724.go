package main

func main() {

	pivotIndex([]int{1, 7, 3, 6, 5, 6})
}

// Prefix Sum (Optimized Space)
// Time Complexity: O(N)
// Space Complexity: O(1)
func pivotIndex(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0
	for i, num := range nums {
		// Right Sum = Total Sum - Left Sum - nums[i]
		// We want check: Left Sum == Right Sum
		if leftSum == totalSum-leftSum-num {
			return i
		}
		leftSum += num
	}

	return -1
}
