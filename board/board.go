package board

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Color uint

const (
	NONE Color = iota
	BLACK
	WHITE
)

const (
	// Unicode characters for "Large Black Circle" and "Large Circle"
	CROSS       = "+"
	WHITE_STONE = "\u2b24"
	BLACK_STONE = "\u25ef"

	// Standard sizes for a Go board.
	SMALL  = 9
	MEDIUM = 13
	LARGE  = 19
)

var drawings = [...]string{CROSS, BLACK_STONE, WHITE_STONE}

type Board struct {
	grid      [][]Color
	size      int
	blackTurn bool
	playing   bool
}

func NewBoard(n int) *Board {
	// Check for standard board sizes.
	if n != SMALL && n != MEDIUM && n != LARGE {
		fmt.Println("Inputs not valid grid sizes.")
		return nil
	}
	b := make([][]Color, n)
	for i := 0; i < n; i++ {
		b[i] = make([]Color, n)
		for j := 0; j < n; j++ {
			b[i][j] = NONE
		}
	}
	return &Board{grid: b, size: n, blackTurn: true, playing: true}
}

func NewBoardFromArray(array [][]Color, turn bool) *Board {
	if (len(array) != SMALL && len(array) != MEDIUM && len(array) != LARGE) ||
		(len(array) != len(array[0])) {
		fmt.Println("Invalid grid size.")
		return nil
	}
	return &Board{grid: array, size: len(array), blackTurn: turn, playing: true}
}

/* type Player struct {
	black bool
	turn  bool
}*/

func (board *Board) Play() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		board.printAndPrompt()
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			break
		}
		board.interpretAndPlace(scanner.Text())
		time.Sleep(time.Millisecond * 100)
	}
}

func (board *Board) interpretAndPlace(input string) {
	array := strings.Split(input, " ")
	if len(input) < 2 {
		fmt.Println("Please enter a row and a column separated by a space.")
		return
	}
	row, err := strconv.ParseInt(array[0], 10, 0)
	if err != nil {
		fmt.Println("Invalid row entered.")
		return
	}
	col, err := strconv.ParseInt(array[1], 10, 0)
	if err != nil {
		fmt.Println("Invalid col entered.")
		return
	}
	row, col = row-1, col-1
	color := BLACK
	if !board.blackTurn {
		color = WHITE
	}
	board.placeStone(int(row), int(col), color)
}

func (board *Board) placeStone(row int, col int, color Color) {
	if row < 0 || col < 0 ||
		row > board.size || col > board.size {
		fmt.Println("Out of bounds!")
		return
	}
	if board.grid[row][col] != NONE {
		fmt.Println("A stone has already been placed here.")
		return
	}
	if board.isIllegal(row, col, color) {
		fmt.Println("That move is suicidal.")
		return
	}
	// board.checkStones(row, col, color)
	board.blackTurn = !(board.blackTurn)
	board.grid[row][col] = color
}

func (board *Board) checkStones(row int, col int) {
	// Check if other stones are in danger.
}

func (board *Board) isIllegal(row int, col int, color Color) bool {
	// Huge if statements
	surround := true
	that := NONE
	if row > 0 {
		that = board.grid[row-1][col]
		if that == NONE {
			return false
		}
		if that == color {
			surround = board.isIllegal(row-1, col, color)
		}
	}
	if col > 0 {
		that = board.grid[row][col-1]
		if that == NONE {
			return false
		}
		if that == color {
			surround = board.isIllegal(row, col-1, color)
		}
	}
	if row < board.size {
		that = board.grid[row+1][col]
		if that == NONE {
			return false
		}
		if that == color {
			surround = board.isIllegal(row+1, col, color)
		}
	}
	if col < board.size {
		that = board.grid[row][col+1]
		if that == NONE {
			return false
		}
		if that == color {
			surround = board.isIllegal(row, col+1, color)
		}
	}
	return surround
}

func (board *Board) printAndPrompt() {
	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			fmt.Print(" " + drawings[board.grid[i][j]])
		}
		fmt.Println()
	}
	if board.blackTurn {
		fmt.Println("black to play.")
	} else {
		fmt.Println("white to play.")
	}
}
