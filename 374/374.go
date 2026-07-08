package main

func main() {
	guessNumber(20)
}

func guessNumber(n int) int {
	lo, hi := 1, n
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch guess(mid) {
		case 0:
			return mid
		case -1:
			hi = mid - 1
		default:
			lo = mid + 1
		}
	}
	return -1
}

func guess(num int) int {

	if num == 6 {
		return 0
	} else if num > 6 {
		return -1
	} else {
		return 1
	}
}
