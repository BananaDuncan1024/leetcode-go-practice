package main

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}

	removeDuplicates(nums1)
	// print(ans)

}

func removeDuplicates(nums []int) int {

	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j] {
			nums[j+1] = nums[i]
			j++
		}
	}
	return j + 1
}
