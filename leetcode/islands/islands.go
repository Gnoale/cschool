package islands

func numIslands(grid [][]byte) int {

	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

// return the neighbors of the cell in all directions
func getNeighbors(grid [][]byte, row int, col int) [][]int {

	neighbors := [][]int{}
	directions := [][]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	for _, direction := range directions {
		newRow := row + direction[0]
		newCol := col + direction[1]
		if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) && grid[newRow][newCol] == '1' {
			neighbors = append(neighbors, []int{newRow, newCol})
		}
	}
	return neighbors
}

func dfs(grid [][]byte, row int, col int) {
	grid[row][col] = '0'
	neighbors := getNeighbors(grid, row, col)
	for _, neighbor := range neighbors {
		dfs(grid, neighbor[0], neighbor[1])
	}
}
