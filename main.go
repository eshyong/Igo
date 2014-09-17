package main

import "github.com/eshyong/go-bot/igo"

func main() {
	// Test
	/*	grid := make([][]igo.Color, 9)
		for i := 0; i < 9; i++ {
			grid[i] = make([]igo.Color, 9)
			for j := 0; j < 9; j++ {
				grid[i][j] = igo.NONE
			}
		}
		grid[2][0] = igo.BLACK
		grid[1][1] = igo.BLACK
		grid[1][2] = igo.BLACK
		grid[3][1] = igo.BLACK
		grid[3][2] = igo.BLACK
		grid[2][3] = igo.BLACK
		grid[2][1] = igo.WHITE
		grid[2][2] = igo.WHITE */
	game := igo.NewGame(igo.SMALL)
	if game != nil {
		game.Play()
	}
}
