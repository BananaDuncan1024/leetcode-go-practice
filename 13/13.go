package main

func romanToInt(s string) int {

	romanMap := map[rune]int{

		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var sum int
	var prev rune
	for _, value := range s {
		sum += romanMap[value]
		if romanMap[prev] < romanMap[value] {
			sum -= romanMap[prev] * 2
		}
		prev = value
	}

	return sum
}

func main() {

	s := "MCMXCIV"

	romanToInt(s)
}
