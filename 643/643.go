package main

func main() {
	nums, k := []int{1, 12, -5, -6, 50, 3}, 4

	findMaxAverage(nums, k)
}

func findMaxAverage(nums []int, k int) float64 {

	i := 0
	j := k - 1

	var latest float64 = -797693134862315708145274237

	for j < len(nums) {

		sum := 0
		for x := i; x <= j; x++ {
			sum += nums[x]
		}

		if float64(sum) > latest {
			latest = float64(sum)
		}
		i++
		j++
	}

	return float64(latest) / float64(k)
}

func findMaxAverage2(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	curr := float64(sum) / float64(k)

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]

		curr = max(curr, float64(sum)/float64(k))
	}

	return curr

}

func max(i, j float64) float64 {
	if i > j {
		return i
	}
	return j
}
