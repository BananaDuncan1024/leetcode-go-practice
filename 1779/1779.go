package main

import (
	"fmt"
	"math"
)

func main() {

	testcase := [][]int{
		{1, 2},
		{3, 1},
		{2, 4},
		{2, 3},
		{4, 4},
	}

	ans := nearestValidPoint(3, 4, testcase)
	print(ans)
}

func nearestValidPoint(x int, y int, points [][]int) int {

	valuePoints := [][]int{}

	distance := float64(^uint(0) >> 1)
	var c float64
	// counter
	point := -1

	fmt.Print(points)

	for _, value := range points {
		// fmt.Print(value)
		if value[0] == x || value[1] == y {
			// fmt.Print(value)
			valuePoints = append(valuePoints, value)
		}
	}

	if len(valuePoints) == 0 {
		return -1
	} else {
		for index, value := range valuePoints {
			c = math.Abs(float64(x-value[0])) + math.Abs(float64(y-value[1]))
			// fmt.Print(c)
			if c < distance {
				distance = c
				point = index
			}
		}
	}

	return point

}
