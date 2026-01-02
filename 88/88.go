package main

import "sort"

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge2(nums1, 3, nums2, 3)
	// print(ans)

}

// brute force - merge and sort
func merge(nums1 []int, m int, nums2 []int, n int) {
	// Merge nums1 and nums2
	nums1 = append(nums1[:n], nums2...)
	// Sort nums1
	sort.Ints(nums1)
}

// Insert One Each Time
func merge2(nums1 []int, m int, nums2 []int, n int) {

	for _, valuenums2 := range nums2 {
		for _, valuenums1 := range nums1 {
			if valuenums2 >= valuenums1 {
				nums1 = append(nums1, valuenums2)
				// fmt.Print(nums1)
			}
		}

	}
}

// Two Pointer
func merge3(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 || n > 0 {
		if n == 0 {
			break
		}
		if m == 0 {
			nums1[n-1] = nums2[n-1]
			n--
			continue
		}
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

// insert for
func merge4(nums1 []int, m int, nums2 []int, n int) {
	i1, i2 := m-1, n-1
	for i := len(nums1) - 1; i >= 0; i-- {
		if i2 < 0 || (i1 >= 0 && nums1[i1] > nums2[i2]) {
			nums1[i] = nums1[i1]
			i1--
		} else {
			nums1[i] = nums2[i2]
			i2--
		}
	}
}
