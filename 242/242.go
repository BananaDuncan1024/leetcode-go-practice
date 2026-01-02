package main

func main() {
	s, t := "anagram", "nagaram"
	isAnagram(s, t)
}

func isAnagram(s string, t string) bool {

	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, value := range s {
		sMap[value]++
	}

	for _, value := range t {
		tMap[value]++
	}

	for key, value := range sMap {
		if tMap[key] != value {
			return false
		}
	}

	return true
}

// O(nlogn)
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var arr [26]int

	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}
