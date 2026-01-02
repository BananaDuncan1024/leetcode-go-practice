package main

import ()

// func canConstruct(ransomNote string, magazine string) bool {

// 	if strings.Contains(magazine, ransomNote) {
// 		return true
// 	} else {
// 		return false
// 	}

// }

// Time: O(n), Space: O(1)
func canConstruct2(ransomNote string, magazine string) bool {
	counts := [26]int{}
	for i := 0; i < len(magazine); i++ {
		counts[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		counts[ransomNote[i]-'a']--
		if counts[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {

	ransomNote := "aa"
	magazine := "aab"

	canConstruct2(ransomNote, magazine)

}
