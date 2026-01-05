package main

import "fmt"

func main() {
	nums1 := []int{2, 3, -2, 4}
	fmt.Printf("Input: %v, Max Product: %d\n", nums1, maxProduct(nums1))

	nums2 := []int{-2, 0, -1}
	fmt.Printf("Input: %v, Max Product: %d\n", nums2, maxProduct(nums2))
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSoFar := nums[0]
	minSoFar := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		curr := nums[i]

		// 關鍵步驟：如果當前數字是負數，
		// 那麼 "最大值 * 負數" 會變成最小值，"最小值 * 負數" 可能變成最大值。
		// 所以我們先交換 maxSoFar 和 minSoFar。
		if curr < 0 {
			maxSoFar, minSoFar = minSoFar, maxSoFar
		}

		// 動態規劃狀態轉移：
		// maxSoFar 是 "當前數字" 與 "當前數字 * 之前的 max" 之間的較大者
		// minSoFar 是 "當前數字" 與 "當前數字 * 之前的 min" 之間的較小者
		// 注意：我們不需要像 Sum 那樣判斷 < 0 重置，因為這裡是乘法，狀態是連續的
		maxSoFar = max(curr, maxSoFar*curr)
		minSoFar = min(curr, minSoFar*curr)

		// 更新全域結果
		result = max(result, maxSoFar)
	}

	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
