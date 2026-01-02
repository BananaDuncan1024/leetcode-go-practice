package main

func main() {

	nums := []int{7, 1, 5, 3, 6, 4}
	maxProfit(nums)

}

func maxProfit(prices []int) int {

	profit := 0
	lowest := prices[0]

	for _, value := range prices {

		if value < lowest {
			lowest = value
		} else if value-lowest > profit {
			profit = value - lowest
		}
	}

	return profit
}
