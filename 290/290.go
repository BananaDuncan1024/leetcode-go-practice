package main

import "strings"

func main() {
	pattern, s := "abba", "dog cat cat dog"

	wordPattern(pattern, s)
}

func wordPattern(pattern string, s string) bool {

	strSlice := strings.Split(s, " ") // "dog cat cat fish"

	if len(pattern) != len(strSlice) { // 1. edge case 處理
		return false
	}

	strMap := make(map[string]byte, len(s))           // 2. dog -> a
	patternMap := make(map[byte]string, len(pattern)) // 2. a -> dog

	for i, sliceAnimal := range strSlice { // 3. 遍歷 strSlice 或是 pattern 都可以，因為長度相同

		if singlePattern, exist := strMap[sliceAnimal]; !exist {
			strMap[sliceAnimal] = pattern[i]
		} else {
			if singlePattern != pattern[i] { // 4. 關鍵，如果已經存在，那看存在的 value 跟 另一張表中的是否相同ㄋ
				return false
			}
		}

		if animal, exist := patternMap[pattern[i]]; !exist {
			patternMap[pattern[i]] = sliceAnimal
		} else {
			if animal != sliceAnimal {
				return false
			}
		}

	}

	return true
}

// better memory usage
func wordPattern2(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(words) != len(pattern) {
		return false
	}

	var ps [26]string
	sp := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		x := ps[pattern[i]-'a']
		y, ok := sp[words[i]]

		if x == "" && !ok {
			ps[pattern[i]-'a'] = words[i]
			sp[words[i]] = pattern[i]
		} else if x != words[i] || y != pattern[i] {
			return false
		}
	}
	return true
}
