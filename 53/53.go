package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Printf("Input: %v\n", nums)
	fmt.Printf("Kadane's Algorithm: %d\n", maxSubArray(nums))
	fmt.Printf("Divide and Conquer: %d\n", maxSubArrayDC(nums))
}

// Kadane's Algorithm (O(n))
func maxSubArray(nums []int) int {
	sum, maxSum := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		maxSum = max(maxSum, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}

// Divide and Conquer Approach (O(n log n))
func maxSubArrayDC(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) / 2

	// Maximum subarray sum in the left half
	leftMax := helper(nums, left, mid)
	// Maximum subarray sum in the right half
	rightMax := helper(nums, mid+1, right)
	// Maximum subarray sum that crosses the midpoint
	crossMax := maxCrossingSum(nums, left, mid, right)

	return max(max(leftMax, rightMax), crossMax)
}

func maxCrossingSum(nums []int, left, mid, right int) int {
	// Find the maximum sum starting from mid and moving left
	leftSum := math.MinInt
	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	// Find the maximum sum starting from mid+1 and moving right
	rightSum := math.MinInt
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return leftSum + rightSum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}