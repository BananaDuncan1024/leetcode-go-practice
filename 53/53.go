package main

func main() {
	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}

// Kadane's Algorithm，這是解決最大子陣列和問題的經典方法
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

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
