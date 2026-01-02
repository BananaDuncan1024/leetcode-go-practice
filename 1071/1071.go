package main

func main() {
	str1, str2 := "ABCABC", "ABC"

	gcdOfStrings(str1, str2)
}

// Complexity
// Time complexity:
// O(n), since we need to compare the strings and concate the two strings. Time complexity is linear.

// Space complexity:
// O(1)
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	len := gcd(len(str1), len(str2))
	return str1[:len]
}

func gcd(a int, b int) int {

	for b != 0 {
		a, b = b, a%b
	}

	return a
}
