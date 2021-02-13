package solution

import (
	"container/list"
)

var directions = [...][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

// TODO: review solution
func shortestPathBinaryMatrix(grid [][]int) int {
	m, n := len(grid)-1, len(grid[0])-1
	if grid[0][0] == 1 || grid[n][m] == 1 {
		return -1
	}
	q := list.New()
	q.PushBack([]int{0, 0, 1})
	for q.Len() != 0 {
		// worst pop() you've ever seen
		a, _ := q.Front().Value.([]int), q.Remove(q.Front())
		if a[0] == n && a[1] == m {
			return a[2]
		}
		for _, v := range directions {
			x, y := a[0]+v[0], a[1]+v[1]
			if x >= 0 && x <= n && y >= 0 && y <= m && grid[x][y] == 0 {
				grid[x][y] = 1
				q.PushBack([]int{x, y, a[2] + 1})
			}
		}
	}
	return -1
}

// func shortestPathBinaryMatrix(grid [][]int) int {
//     if grid[0][0] == 1 {
// 		return -1
// 	}

// 	row, col := len(grid), len(grid[0])
// 	if row == 1 && col == 1 {
// 		return 1
// 	}

// 	dx, dy := []int{-1, 0, 1, 1, 1, 0, -1, -1}, []int{-1, -1, -1, 0, 1, 1, 1, 0}
// 	queue := []int{grid[0][0]}
//  grid[0][0] = 1

// 	for len(queue) > 0 {
// 		cur := queue[0]
// 		queue = queue[1:]
// 		x, y := cur/col, cur%col

// 		for i := range dx {
// 			nx, ny := x+dx[i], y+dy[i]
// 			if nx >= 0 && nx < row && ny >= 0 && ny < col && grid[nx][ny] == 0 {
// 				queue = append(queue, nx*col+ny)
// 				grid[nx][ny] = grid[x][y] + 1
// 				if nx == row-1 && ny == col-1 {
// 					return grid[nx][ny]
// 				}
// 			}
// 		}
// 	}
// 	return -1
// }
