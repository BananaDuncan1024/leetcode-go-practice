package main

import "strings"

func main() {

	s := "hello"
	reverseVowels(s)

}
func reverseVowels(s string) string {

	started := 0
	ended := len(s) - 1
	vowels := "aeiouAEIOU"
	runeSlice := []rune(s)

	for started < ended {

		for started < ended && !strings.Contains(vowels, string(runeSlice[started])) {
			started++
		}
		// Move the `high` pointer to the left until it points to a vowel
		for started < ended && !strings.Contains(vowels, string(runeSlice[ended])) {
			ended--
		}

		if started < ended {
			// Swap the vowels at positions pointed by `low` and `high`
			runeSlice[started], runeSlice[ended] = runeSlice[ended], runeSlice[started]
			started++
			ended--
		}

	}

	return string(runeSlice)

}
