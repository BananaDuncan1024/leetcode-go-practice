package main

func main() {

	word1, word2 := "abc", "pqr"

	mergeAlternately1(word1, word2)

}

// O(m+n)
// O(m+n)
func mergeAlternately1(word1 string, word2 string) string {

	length := len(word1) + len(word2)

	var str string

	for i := 0; i < length; i++ {
		if i < len(word1) {
			str += string(word1[i])
		}
		if i < len(word2) {
			str += string(word2[i])
		}
	}

	return str
}

// O(m+n)
func mergeAlternately2s(word1 string, word2 string) string {

	length := len(word1)

	var str string

	if len(word2) < len(word1) {
		length = len(word2)
	}

	//merge upto the shorter stringlength
	for i := 0; i < length; i++ {
		str += string(word1[i]) + string(word2[i])
	}

	//add the remaining chars from the longer string
	str += word1[length:] + word2[length:]

	return str
}
