package main

func main() {

	findMiddleIndex([]int{2, 3, -1, 8, 4})
}

func findMiddleIndex(nums []int) int {

	totalsum := 0
	for _, n := range nums {
		totalsum += n
	}

	leftsum := 0

	// We want check: Left Sum == Right Sum
	for i, num := range nums {
		if leftsum == totalsum-leftsum-num {
			return i
		}
		leftsum += num
	}

	return -1
}
