package pointer

import "fmt"

func Print2DBoard(board [][]string) { 
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("%s ", board[i][j])
		}
		fmt.Println()
	}
}

func CheckWinner(board [][]string) string {
	// Check rows
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != "_" {
			return board[i][0]
		}
	}

	// Check columns
	for j := 0; j < 3; j++ {
		if board[0][j] == board[1][j] && board[1][j] == board[2][j] && board[0][j] != "_" {
			return board[0][j]
		}
	}

	// Check diagonals
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != "_" {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != "_" {
		return board[0][2]
	}

	return ""
}

func TicTacToe() { 
	var board = [][]string {
		[]string {"_", "_", "_"},
		[]string {"_", "_", "_"},
		[]string {"_", "_", "_"},
	}
	Print2DBoard(board)

	// player takes turn
	board[0][0] = "X"
	board[0][2] = "X"
	board[0][1] = "X"
	board[1][0] = "O"
	board[1][1] = "O"
	board[1][2] = "X"
	board[2][0] = "O"
	board[2][1] = "X"
	board[2][2] = "O"

	winner := CheckWinner(board)

	Print2DBoard(board);
	if winner != "" {
		fmt.Printf("Winner is: %s\n", winner)
	} else {
		fmt.Println("Game is not over yet.")
	}
}
