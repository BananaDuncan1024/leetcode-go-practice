package main

// n by n
func numberOfSteps(num int) int {

	var step int

	var countSteps func(num int)

	countSteps = func(num int) {
		if num == 0 {
			return
		} else if num%2 == 0 {
			countSteps(num / 2)
		} else {
			countSteps(num - 1)
		}

		step++
	}
	countSteps(num)

	return step

}

// n by 1
func numberOfSteps2(num int) int {
	var steps int

	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}

		steps++
	}

	return steps
}

// logn
func numberOfStep3(num int) int {
	if num == 0 {
		return 0
	}

	result := 0
	for num > 0 {
		if num&1 == 1 {
			result += 2
		} else {
			result += 1
		}

		num >>= 1
	}

	return result - 1
}

func main() {

	numberOfSteps(14)

}
