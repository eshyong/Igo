package igo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Color uint
type Direction uint

const (
	NONE Color = iota
	BLACK
	WHITE
)

const (
	START Direction = iota
	NORTH
	EAST
	SOUTH
	WEST
)

const (
	// Unicode characters for "Large Black Circle" and "Large Circle"
	CROSS       = "+"
	WHITE_STONE = "\u2b24"
	BLACK_STONE = "\u25ef"

	// Standard sizes for a Go game.
	SMALL  = 9
	MEDIUM = 13
	LARGE  = 19
)

var drawings = [...]string{CROSS, BLACK_STONE, WHITE_STONE}

type Game struct {
	board     [][]Color
	size      int
	blackTurn bool
	playing   bool
}

type Point struct {
	row int
	col int
}

func newPoint(row int, col int) *Point {
	return &Point{row: row, col: col}
}

func NewGame(n int) *Game {
	// Check for standard game sizes.
	if n != SMALL && n != MEDIUM && n != LARGE {
		fmt.Println("Inputs not valid board sizes.")
		return nil
	}
	b := make([][]Color, n)
	for i := 0; i < n; i++ {
		b[i] = make([]Color, n)
		for j := 0; j < n; j++ {
			b[i][j] = NONE
		}
	}
	return &Game{board: b, size: n, blackTurn: true, playing: true}
}

func NewGameFromArray(array [][]Color, turn bool) *Game {
	if (len(array) != SMALL && len(array) != MEDIUM && len(array) != LARGE) ||
		(len(array) != len(array[0])) {
		fmt.Println("Invalid board size.")
		return nil
	}
	return &Game{board: array, size: len(array), blackTurn: turn, playing: true}
}

/* type Player struct {
	black bool
	turn  bool
}*/

func (game *Game) Play() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		game.PrintAndPrompt()
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			break
		}
		game.interpretAndPlace(scanner.Text())
		time.Sleep(time.Millisecond * 100)
	}
}

func (game *Game) interpretAndPlace(input string) {
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
	if !game.blackTurn {
		color = WHITE
	}
	game.placeStone(int(row), int(col), color)
}

func (game *Game) placeStone(row int, col int, color Color) {
	if row < 0 || col < 0 ||
		row > game.size || col > game.size {
		fmt.Println("Out of bounds!")
		return
	}
	if game.board[row][col] != NONE {
		fmt.Println("A stone has already been placed here.")
		return
	}
	if game.IsDead(row, col, color, START) {
		fmt.Println("That move is suicidal.")
		return
	}
	// Update game.
	game.board[row][col] = color
	game.CheckSurroundingStones(row, col, color)

	game.blackTurn = !(game.blackTurn)
	game.board[row][col] = color
}

func (game *Game) CheckSurroundingStones(row int, col int, color Color) {
	// Check if enemy groups are dead.
	game.RemoveIfDead(row-1, col, color)
	game.RemoveIfDead(row+1, col, color)
	game.RemoveIfDead(row, col-1, color)
	game.RemoveIfDead(row, col+1, color)
}

func (game *Game) RemoveIfDead(row int, col int, color Color) {
	if row < 0 || row >= game.size || col < 0 || col >= game.size || game.board[row][col] == color {
		return
	}
	var other Color
	if color == BLACK {
		other = WHITE
	} else {
		other = BLACK
	}
	if game.IsDead(row, col, other, START) {
		game.removeRecursively(row, col, other)
	}
}

func (game *Game) removeRecursively(row int, col int, color Color) {
	if row < 0 || row >= game.size || col < 0 || col >= game.size || color != game.board[row][col] {
		return
	}
	game.board[row][col] = NONE
	// Optimize this later
	game.removeRecursively(row-1, col, color)
	game.removeRecursively(row+1, col, color)
	game.removeRecursively(row, col+1, color)
	game.removeRecursively(row, col-1, color)
}

func (game *Game) IsDead(row int, col int, color Color, from Direction) bool {
	// Stack iteration instead of recursion
	// Create a map to keep track of added rows
	stack := make([]*Point, 0, game.size*game.size)
	visited := make(map[string]bool)

	curr := newPoint(row, col)
	hash := stringHash(row, col)
	visited[hash] = true
	stack = append(stack, curr)
	for len(stack) > 0 {
		// Pop point off stack
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		currRow, currCol := curr.row, curr.col
		hash = stringHash(currRow, currCol)
		visited[hash] = true

		if game.board[currRow][currCol] == NONE {
			// Found a liberty!
			return false
		} else if game.board[currRow][currCol] == color {
			// Check left, right, top, and bottom.
			// We don't want to add points that we already traversed.
			if currRow > 0 {
				hash = stringHash(currRow-1, currCol)
				if !visited[hash] {
					stack = append(stack, newPoint(currRow-1, currCol))
				}
			}
			if currRow < game.size-1 {
				hash = stringHash(currRow+1, currCol)
				if !visited[hash] {
					stack = append(stack, newPoint(currRow+1, currCol))
				}
			}
			if currCol > 0 {
				hash = stringHash(currRow, currCol-1)
				if !visited[hash] {
					stack = append(stack, newPoint(currRow, currCol-1))
				}
			}
			if currCol < game.size-1 {
				hash = stringHash(currRow, currCol+1)
				if !visited[hash] {
					stack = append(stack, newPoint(currRow, currCol+1))
				}
			}
		}
		// Ignore different colored stones
	}
	return true
}

func stringHash(row int, col int) string {
	return strconv.FormatInt(int64(row), 10) + ", " + strconv.FormatInt(int64(col), 10)
}

func PrintColor(color Color) {
	switch color {
	case BLACK:
		fmt.Println("BLACK")
	case WHITE:
		fmt.Println("WHITE")
	case NONE:
		fmt.Println("EMPTY")
	}
}

func (game *Game) PrintAndPrompt() {
	for i := 0; i < game.size; i++ {
		for j := 0; j < game.size; j++ {
			fmt.Print(" " + drawings[game.board[i][j]])
		}
		fmt.Println()
	}
	if game.blackTurn {
		fmt.Println("black to play.")
	} else {
		fmt.Println("white to play.")
	}
}
