package main

import "github.com/eshyong/go-bot/igo"

func test() {
	// Testing
	grid := make([][]igo.Color, 9)
	for i := 0; i < 9; i++ {
		grid[i] = make([]igo.Color, 9)
		for j := 0; j < 9; j++ {
			grid[i][j] = igo.NONE
		}
	}
	grid[0][5] = igo.WHITE
	grid[1][5] = igo.WHITE
	grid[2][0] = igo.WHITE
	grid[2][1] = igo.WHITE
	grid[2][2] = igo.WHITE
	grid[2][3] = igo.WHITE
	grid[2][4] = igo.WHITE

	grid[0][0] = igo.BLACK
	grid[0][2] = igo.BLACK
	grid[0][4] = igo.BLACK
	grid[1][0] = igo.BLACK
	grid[1][1] = igo.BLACK
	grid[1][2] = igo.BLACK
	grid[1][3] = igo.BLACK
	grid[1][4] = igo.BLACK

	game := igo.NewGameFromArray(grid, false)
	game.Play()
}

func main() {
	test()
	/*game := igo.NewGame(igo.MEDIUM)
	if game != nil {
		game.Play()
	}*/
}
