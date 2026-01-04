package main

func main() {

	commonFactors(6, 12)
}

func commonFactors(a int, b int) int {

	max := a
	if b < a {
		max = b
	}
	count := 0

	for i := 1; i <= max; i++ {
		if a%i == 0 && b%i == 0 {
			count++
		}

	}

	return count
}

func gcd(a, b int) int {

	if b > a {
		a, b = b, a
	}

	for b != 0 {

		r := a % b
		a, b = b, r
	}

	return a
}

// 更低的時間複雜度
func commonFactors2(a int, b int) int {

	gcd := gcd(a, b)

	counter := 0

	for i := 1; i <= gcd; i++ {

		if gcd%i == 0 {
			counter += 1
		}
	}

	return counter
}
