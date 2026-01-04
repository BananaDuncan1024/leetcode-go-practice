package main

func main() {

	nums := []int{7, 1, 5, 3, 6, 4}
	maxProfit(nums)

}

// O (n^2) 過久
func maxProfit2(prices []int) int {

	max := 0

	for index, price := range prices {
		for j := index + 1; j < len(prices); j++ {
			if prices[j]-price > max {
				max = prices[j] - price
			}
		}
	}
	return max
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
