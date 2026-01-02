package main

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	kidsWithCandies(candies, extraCandies)
}

// Time complexity: O(n)
// Space complexity: O(n)
func kidsWithCandies(candies []int, extraCandies int) []bool {

	maxKid := 0

	output := make([]bool, len(candies))

	for _, value := range candies {
		if value > maxKid {
			maxKid = value
		}
	}

	for index, value := range candies {
		if value+extraCandies >= maxKid {
			output[index] = true
		} else {
			output[index] = false
		}
	}

	return output
}

func kidsWithCandies2(candies []int, extraCandies int) []bool {
	return nil
}
