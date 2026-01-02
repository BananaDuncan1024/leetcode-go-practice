package main

func main() {

	test := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},

		// []byte{"1", "1", "0", "1", "0"},
		// []byte{"1", "1", "0", "0", "0"},
		// []byte{"0", "0", "0", "0", "0"},
	}

	// test := [][]int{1, 2, 4, 8, 16, 32, 64, 128}

	// var t string

	// t = "t"

	// [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]];
	// test := [][]byte("11110", "11010")
	// fmt.Printf("%v", test)
	count := numIslands(test)
	print(count)

}

func numIslands(grid [][]byte) int {
	numOfIslands := 0
	// Initialize 2D matrix.
	visited := make([][]bool, len(grid))

	// print(len(grid))

	// print(len(grid[0]))

	for row := range visited {
		visited[row] = make([]bool, len(grid[0]))
	}

	// For each cell - we perform dfs.
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			// Skip performing dfs from water.
			if grid[x][y] == '0' {
				continue
			}
			// We mark the board as visited as we dfs on it.
			if visited[x][y] {
				continue
			}
			// DFS will mark all connected land from this cell as visited.
			islandDFS(grid, visited, x, y)
			numOfIslands++
		}
	}

	return numOfIslands
}

func islandDFS(grid [][]byte, visited [][]bool, x, y int) {
	// Boundary check.
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return
	}
	// Return if we've already visited it.
	if visited[x][y] {
		return
	}
	// Return if we hit water.
	if grid[x][y] == '0' {
		return
	}

	// Mark current cell as visited.
	visited[x][y] = true

	// A neighbor can be traversed to (top, bottom, right, left).
	for _, direction := range getDirections() {
		dx, dy := direction[0], direction[1]
		islandDFS(grid, visited, x+dx, y+dy)
	}
}

func getDirections() [][]int {
	return [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
}
