package main

// Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

// answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
// answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
// Note that the integers in the lists may be returned in any order.
func main() {
	findDifference([]int{1, 2, 3}, []int{2, 4, 6})
}

// Time complexity: O(n + m)
// Space complexity: O(n + m)
func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := map[int]bool{}
	map2 := map[int]bool{}
	for _, n := range nums1 {
		map1[n] = true
	}
	for _, n := range nums2 {
		map2[n] = true
	}

	diff1, diff2 := []int{}, []int{}
	for k := range map1 {
		if !map2[k] {
			diff1 = append(diff1, k)
		}
	}
	for k := range map2 {
		if !map1[k] {
			diff2 = append(diff2, k)
		}
	}

	return [][]int{diff1, diff2}
}
