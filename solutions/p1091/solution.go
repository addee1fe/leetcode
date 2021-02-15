package solution

// var dx = [...]int{-1, 0, 1, 1, 1, 0, -1, -1}
// var dy = [...]int{-1, -1, -1, 0, 1, 1, 1, 0}

// func shortestPathBinaryMatrix(grid [][]int) int {
// 	m, n := len(grid)-1, len(grid[0])-1
// 	if grid[0][0] == 1 || grid[n][m] == 1 {
// 		return -1
// 	}
// 	q := []int{grid[0][0]}
// 	grid[0][0] = 1
// 	for len(q) > 0 {
// 		tmp := q[0]
// 		q = q[1:]
// 		x, y := tmp/m, tmp%m
// 		for i := range dx {
// 			nx, ny := x+dx[i], y+dy[i]
// 			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 0 {
// 				q = append(q, nx*n+ny)
// 				grid[nx][ny] = grid[x][y] + 1
// 				if nx == m-1 && ny == n-1 {
// 					return grid[nx][ny]
// 				}
// 			}
// 		}
// 	}
// 	return -1
// }

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return 1
	}
	dx := [...]int{-1, 0, 1, 1, 1, 0, -1, -1}
	dy := [...]int{-1, -1, -1, 0, 1, 1, 1, 0}
	queue := []int{grid[0][0]}
	grid[0][0] = 1

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		x, y := cur/n, cur%n
		for i := range dx {
			nx, ny := x+dx[i], y+dy[i]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 0 {
				queue = append(queue, nx*n+ny)
				grid[nx][ny] = grid[x][y] + 1
				if nx == m-1 && ny == n-1 {
					return grid[nx][ny]
				}
			}
		}
	}
	return -1
}
