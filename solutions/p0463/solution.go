package solution

func islandPerimeter(grid [][]int) int {
	var p int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				p += 4
				if i > 0 && grid[i-1][j] == 1 {
					p -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					p -= 2
				}
			}
		}
	}
	return p
}
