package islandcount

func numIslands(grid [][]int) int {
	countIsland := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if int(grid[i][j]) == 1 {
				countIsland++
				destroyIsland(grid, i, j)
			}
		}
	}
	return countIsland
}

func destroyIsland(grid [][]int, i, j int) {
	width := len(grid[0])
	height := len(grid)
	if i >= height || i < 0 {
		return
	}
	if j >= width || j < 0 {
		return
	}
	if grid[i][j] != 1 {
		return
	}
	grid[i][j] = 0
	destroyIsland(grid, i+1, j)
	destroyIsland(grid, i, j+1)
	destroyIsland(grid, i-1, j)
	destroyIsland(grid, i, j-1)
	return
}
