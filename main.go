package main

import "github.com/eshyong/go-bot/board"

func main() {
	b := board.NewBoard(board.SMALL)
	if b != nil {
		b.Play()
	}
}
