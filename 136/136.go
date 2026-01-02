package main

func main() {
	nums := []int{2, 2, 1}
	singleNumber(nums)
}

func singleNumber(nums []int) int {

	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}
	return 1
}

func singleNumber2(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}

	return result
}
