package solution

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	n, m, c := len(grid), len(grid[0]), 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				dfs(grid, i, j, m, n)
				c++
			}
		}
	}
	return c
}

func dfs(grid [][]byte, i, j, m, n int) {
	if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i+1, j, m, n)
	dfs(grid, i-1, j, m, n)
	dfs(grid, i, j+1, m, n)
	dfs(grid, i, j-1, m, n)
}
