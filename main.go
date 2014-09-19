package main

import "github.com/eshyong/go-bot/igo"

func main() {
	// Test
	/*grid := make([][]igo.Color, 9)
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
	grid[2][2] = igo.WHITE

	grid[6][8] = igo.BLACK
	grid[6][7] = igo.BLACK
	grid[7][6] = igo.BLACK
	grid[8][6] = igo.BLACK
	grid[7][7] = igo.WHITE
	grid[7][8] = igo.WHITE
	grid[8][7] = igo.WHITE
	grid[8][8] = igo.WHITE

	game := igo.NewGameFromArray(grid, true)
	game.PrintAndPrompt()
	fmt.Println(game.IsDead(2, 1, igo.WHITE, igo.START))
	fmt.Println(game.IsDead(8, 8, igo.WHITE, igo.START))*/
	game := igo.NewGame(igo.SMALL)
	if game != nil {
		game.Play()
	}
}
