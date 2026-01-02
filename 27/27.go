package main

func main() {

	nums1 := []int{3, 2, 2, 3}

	removeElement(nums1, 3)
	// print(ans)

}

// O(n) - Two Pointer
func removeElement(nums []int, val int) int {

	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

	return j
}

// Read/Write Pointer (Fast-Slow Pointer)
func removeElement2(nums []int, val int) int {

	ptr1, ptr2 := 0, len(nums)-1

	for ptr1 <= ptr2 {

		if nums[ptr1] != val {
			ptr1++
		} else {
			nums[ptr1], nums[ptr2] = nums[ptr2], nums[ptr1]
			ptr2--
		}

	}

	return ptr1
}

func removeElement3(nums []int, val int) int {

	index := 0

	for _, num := range nums {
		if num != val {
			nums[index] = num
			index++
		}
	}

	return index
}
