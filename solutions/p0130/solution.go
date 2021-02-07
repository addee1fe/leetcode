package solution

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])

	for i := 0; i < m; i++ {
		find(board, i, 0)
		find(board, i, n-1)
	}
	for j := 1; j < n-1; j++ {
		find(board, 0, j)
		find(board, m-1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'Z' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func find(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return
	}
	if board[x][y] != 'O' {
		return
	}
	board[x][y] = 'Z'
	find(board, x-1, y)
	find(board, x+1, y)
	find(board, x, y-1)
	find(board, x, y+1)
}
