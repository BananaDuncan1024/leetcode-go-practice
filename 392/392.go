package main

func main() {
	s, t := "abc", "ahbgdc"
	isSubsequence(s, t)

}

func isSubsequence(s string, t string) bool {

	count := 0

	if len(s) == 0 {
		return true
	}

	for _, value := range t {
		if rune(s[count]) == value {
			count++
		}
		if count == len(s) {
			return true
		}
	}

	return false
}
