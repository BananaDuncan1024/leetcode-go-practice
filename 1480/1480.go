package main

func runningSum(nums []int) []int {

	temp := 0
	var ans []int

	if len(nums) == 0 {
		return nil
	}
	for _, value := range nums {

		temp += value
		ans = append(ans, temp)
	}

	return ans
}

func main() {

	nums := []int{1, 1, 1, 1, 1}
	runningSum(nums)

}
