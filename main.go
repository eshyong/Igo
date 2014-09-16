package main

import "github.com/eshyong/go-bot/board"

func main() {
	grid := make([][]board.Color, 9)
	for i := 0; i < 9; i++ {
		grid[i] = make([]board.Color, 9)
		for j := 0; j < 9; j++ {
			grid[i][j] = board.NONE
		}
	}
	grid[2][2] = board.BLACK
	grid[2][0] = board.BLACK
	grid[1][1] = board.BLACK
	grid[3][1] = board.BLACK
	b := board.NewBoardFromArray(grid, false)
	if b != nil {
		b.Play()
	}
}
