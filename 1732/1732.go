package main

func main() {
	gain := []int{-5, 1, 5, 0, -7}
	largestAltitude(gain)
}

// prefix sum

func largestAltitude(gain []int) int {

	altitudes := make([]int, len(gain)+1)

	altitudes[0] = 0 // initial altitude

	max := 0

	for i := 0; i < len(gain); i++ {
		altitudes[i+1] = altitudes[i] + gain[i]

		if altitudes[i+1] > max {
			max = altitudes[i+1]
		}
	}

	return max
}

func largestAltitude2(gain []int) int {
	// after checking other solutions, I realized
	// I don't need an array for all altitudes,
	// I just need the max
	last := 0
	max := last
	for _, v := range gain {
		new := last + v
		last = new
		if new > max {
			max = new
		}
	}
	return max
}
