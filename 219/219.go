package main

import "math"

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3

	containsNearbyDuplicate(nums, k)
}

func containsNearbyDuplicate(nums []int, k int) bool {

	m := make(map[int]int)

	for i, v := range nums {
		if _, ok := m[v]; ok {
			if math.Abs(float64(i-m[v])) <= float64(k) {
				return true
			}
		}
		m[v] = i

	}

	return false
}
