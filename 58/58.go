package main

import "strings"

func main() {

	s := "Hello World"
	lengthOfLastWord(s)

}

// from the latest
func lengthOfLastWord(s string) int {

	count := 0
	for i := len(s); i > 0; i-- {
		if s[i-1:i] != " " {
			count += 1
		} else if s[i-1:i] == " " && count > 0 {
			return count
		}
	}
	return count
}

// not the best runtime
func lengthOfLastWord2(s string) int {
	s = strings.TrimSpace(s)
	v := strings.Split(s, " ")
	ans := v[len(v)-1]
	return len(ans)
}
