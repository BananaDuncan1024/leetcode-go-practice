package main

func main() {
	n := 19
	isHappy(n)
}

func isHappy(n int) bool {

	m := make(map[int]struct{})

	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		}

		m[n] = struct{}{}
		n = calculateSquaresSum(n)
	}
	return n == 1
}

func calculateSquaresSum(n int) int {
	sum := 0

	for n != 0 {
		r := n % 10
		sum += r * r

		n = n / 10
	}
	return sum
}
