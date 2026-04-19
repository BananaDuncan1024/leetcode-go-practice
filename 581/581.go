package main

import "slices"

func main() {
	findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15})
}

// 一班解法 O(n log n)
func findUnsortedSubarray(nums []int) int {

	sorted_array := make([]int, len(nums))
	copy(sorted_array, nums)

	// 2. 進行排序 (O(n log n))
	slices.Sort(sorted_array)

	started, ended := 0, len(nums)-1

	// 3. 從頭開始比較，找到第一個不相同的元素位置
	for started < len(nums) && nums[started] == sorted_array[started] {
		started++
	}

	if started == len(nums) {
		return 0
	}

	// 找到最後一個不一樣的地方 (右邊界)
	for ended >= started && nums[ended] == sorted_array[ended] {
		ended--
	}

	return ended - started + 1
}

func findUnsortedSubarrayTwoPointers(nums []int) int {

	left, right := -1, -2

	if len(nums) < 2 {
		return 0
	}

	max, min := nums[0], nums[len(nums)-1]

	for i := 0; i < len(nums); i++ {
		if nums[i] < max {
			right = i
		} else {
			max = nums[i]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > min {
			left = i
		} else {
			min = nums[i]
		}
	}

	return right - left + 1
}
