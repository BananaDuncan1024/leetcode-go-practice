package main

func main() {
	haystack, needle := "sadbutsad", "sad"
	strStr(haystack, needle)

}

// fail
func strStr(haystack string, needle string) int {

	if len(needle) > len(haystack) {
		return -1
	}

	for index, value := range haystack {

		if value == rune(needle[0]) {
			for i := 0; i < len(needle); i++ {
				if rune(needle[i]) != rune(haystack[index+i]) {
					break
				}
				if i == len(needle)-1 {
					return index
				}
			}
		}
	}

	return -1
}

// Time complexity: O(N)
// Space complexity: O(1)
func strStr2(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func strStr3(haystack string, needle string) int {
	return -1
}
