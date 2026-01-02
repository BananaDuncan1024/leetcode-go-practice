package main

func maximumWealth(accounts [][]int) int {

	var ans int
	var temp int

	for _, value := range accounts {
		temp = 0
		for _, value2 := range value {
			temp += value2
		}
		if temp > ans {
			ans = temp
		}
	}

	return ans

}

func main() {

	accounts := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	maximumWealth(accounts)

}
