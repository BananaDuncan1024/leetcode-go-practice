package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {

	var reversedNum int

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	for x > reversedNum {
		reversedNum = reversedNum*10 + x%10
		x = x / 10
	}

	return x == reversedNum || x == reversedNum/10
}

func main() {

	x := 121
	// y := 123
	binary := strconv.FormatInt(int64(x), 2)
	fmt.Println("Binary:", binary)

	// isPalindrome(x)s

}
