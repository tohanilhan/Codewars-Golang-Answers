package kata

func IsSolved(board [3][3]int) int {
	wins := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
		{0, 4, 8}, {2, 4, 6}, // Diagonals
	}

	// Check if any player won
	for _, row := range wins {
		a, b, c := row[0], row[1], row[2]
		if board[a/3][a%3] != 0 && board[a/3][a%3] == board[b/3][b%3] && board[b/3][b%3] == board[c/3][c%3] {
			return board[a/3][a%3]
		}
	}

	// Check if the game is still ongoing or a draw
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				return -1 // Board is not yet finished
			}
		}
	}

	return 0 // Cat's game (draw)
}
