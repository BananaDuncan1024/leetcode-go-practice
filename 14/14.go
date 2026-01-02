package main

import (
	"strings"
)

// func longestCommonPrefix2(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}
// 	for index, _ := range strs {
// 		for !strings.HasPrefix(strs[index], prefix) {
// 			prefix = prefix[:len(prefix)-1]
// 		}
// 	}
// }

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for index, _ := range strs {
		for !strings.HasPrefix(strs[index], prefix) {
			prefix = prefix[:len(prefix)-1]
		}
	}

	return prefix
}

func main() {

	strs := []string{"flower", "flow", "flight"}

	longestCommonPrefix(strs)

}
