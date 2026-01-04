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

// 空間複雜度 O(1)
func commonFactors3(a int, b int) int {
	m := a
	if b < a {
		m = b
	}

	no_of_factors := 0

	for n := 1; n <= m/2; n++ {
		if a%n == 0 && b%n == 0 {
			no_of_factors++
		}
	}

	// include min(a, b) if it is a common factor
	if a%m == 0 && b%m == 0 {
		no_of_factors++
	}

	return no_of_factors
}
