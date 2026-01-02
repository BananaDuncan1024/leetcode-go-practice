package main

func main() {
	s := "abcabcbb"
	lengthOfLongestSubstring(s)
}

// fail solution
func lengthOfLongestSubstring(s string) int {

	// init temp
	temp := ""

	j := 1

	for _, value := range s {
		if temp == string(value) {
			temp = string(value)
			j = 1

		} else if temp == "" {
			temp = string(value)
		} else {
			j++

		}
	}

	return j
}

// sliding window solution
// https://medium.com/%E6%8A%80%E8%A1%93%E7%AD%86%E8%A8%98/%E6%BC%94%E7%AE%97%E6%B3%95%E7%AD%86%E8%A8%98%E7%B3%BB%E5%88%97-two-pointer-%E8%88%87sliding-window-8742f45f3f55
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}

	res := -1 << 63

	l := 0
	m := make(map[byte]int)
	m[s[0]]++

	for i := 1; i < len(s); i++ {
		m[s[i]]++

		for m[s[i]] > 1 {
			m[s[l]]--

			if m[s[l]] <= 0 {
				delete(m, s[l])
			}

			l++
		}

		if len(m) > res {
			res = len(m)
		}
	}

	if len(m) > res {
		res = len(m)
	}

	return res
}
