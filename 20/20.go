package main

func isValid(s string) bool {

	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	hashmap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []rune{}

	for _, value := range s {

		if _, ok := hashmap[value]; ok {
			stack = append(stack, value)
		} else if len(stack) == 0 || hashmap[stack[len(stack)-1]] != value {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}

	}

	return len(stack) == 0
}

func main() {

	test := "()[]{}"

	isValid(test)

}
