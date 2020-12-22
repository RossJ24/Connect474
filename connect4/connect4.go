package connect4

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/RossJ24/connect474/printing"
)

// Connect4 represents a Connect4 game
type Connect4 struct {
	// Grid layout of the game
	Layout Grid
	// int representing the player whose turn it is
	Turn int
	//  number of remaining pieces for player 1
	Nump1Pieces int
	//  number of remaining pieces for player 2
	Nump2Pieces int
	// pointer to int of player who won the game
	Winner *int
	// positions that player 1 has played
	P1positions []Coordinate
	// positions that player 2 has played
	P2positions []Coordinate
}

// NewConnect4 returns a new Connect 4 game
func NewConnect4() Connect4 {
	return Connect4{
		Turn:        1,
		Nump1Pieces: 21,
		Nump2Pieces: 21,
		P1positions: make([]Coordinate, 0),
		P2positions: make([]Coordinate, 0),
	}
}

//copyPositions copies the coordinates of each player's pieces
func (c4 *Connect4) copyPositions(player int) []Coordinate {
	if player == 1 {
		cpy := make([]Coordinate, len(c4.P1positions))
		copy(cpy, c4.P1positions)
		return cpy
	}
	cpy := make([]Coordinate, len(c4.P2positions))
	copy(cpy, c4.P2positions)
	return cpy
}

//CopyConnect4 makes a deep copy of a Connect4 game
func (c4 *Connect4) CopyConnect4() *Connect4 {
	return &Connect4{
		Turn:        c4.Turn,
		Nump1Pieces: c4.Nump1Pieces,
		Nump2Pieces: c4.Nump2Pieces,
		Layout:      c4.Layout.CopyGrid(),
		P1positions: c4.copyPositions(1),
		P2positions: c4.copyPositions(2),
	}
}

// lastPlayerTurn returns the player who just played
func (c4 *Connect4) lastPlayerTurn() int {
	if c4.Turn == 1 {
		return 2
	}
	return 1

}

// updateStats updates the statics of a Connect4 game
func (c4 *Connect4) updateStats(Row, Col int) {
	if c4.Turn == 1 {
		c4.P1positions = append(c4.P1positions, newCoord(Row, Col))
		c4.Nump1Pieces--
	} else {
		c4.P2positions = append(c4.P2positions, newCoord(Row, Col))
		c4.Nump2Pieces--
	}

}

// GetWinner returns the winner as an int
func (c4 *Connect4) GetWinner() int {
	if c4.Winner == nil {
		return 0
	}
	return *c4.Winner
}

// GetReward returns the reawrd from player 1's POV
func (c4 *Connect4) GetReward() int {
	if c4.Winner == nil {
		return 0
	} else if *c4.Winner == 1 {
		return 1
	}
	return -1
}

// PossibleMoves returns the possible moves from the current state
func (c4 *Connect4) PossibleMoves() []int {
	possibleMoves := make([]int, 0)

	for i, ele := range c4.Layout[0] {
		if ele == 0 {
			possibleMoves = append(possibleMoves, i+1)
		}
	}
	return possibleMoves
}

// Changes the turn of the current player (oscillates between 1 and 2)
func (c4 *Connect4) changeTurn() {
	c4.Turn = c4.Turn%2 + 1
}

// NextTurn return the next player's turn based on the current turn
func NextTurn(turn int) int {
	return turn%2 + 1
}

// Move makes a connect 4 move
func (c4 *Connect4) Move(Column int, print bool) error {
	Column--
	if Column < 0 || Column >= 7 {
		return errors.New("Column index out of range [1...7], try again")
	}
	for i := len(c4.Layout) - 1; i >= 0; i-- {
		if c4.Layout[i][Column] == 0 {
			c4.Layout[i][Column] = c4.Turn
			c4.updateStats(i, Column)
			c4.changeTurn()
			if print {
				printing.Clear()
				c4.Print()
			}
			c4.isWinningMove(c4.lastPlayerTurn())
			return nil
		}
	}
	return errors.New("Column is full. try again")
}

// UndoMove undoes the last move
func (c4 *Connect4) UndoMove() {
	lastmoveplayer := c4.lastPlayerTurn()
	if lastmoveplayer == 1 {
		lastmove := c4.P1positions[len(c4.P1positions)-1]
		c4.P1positions = c4.P1positions[:len(c4.P1positions)-1]
		c4.Nump1Pieces++
		c4.Layout[lastmove.Row][lastmove.Col] = 0
	} else {
		lastmove := c4.P2positions[len(c4.P2positions)-1]
		c4.P2positions = c4.P2positions[:len(c4.P2positions)-1]
		c4.Nump2Pieces++
		c4.Layout[lastmove.Row][lastmove.Col] = 0
	}
	c4.Winner = nil
	c4.changeTurn()
}

// Print prints the game on the console
func (c4 *Connect4) Print() {
	for r, Row := range c4.Layout {
		if r == 0 {
			for i := 0; i < 7; i++ {
				fmt.Print(" " + strconv.Itoa(i+1))
			}
			fmt.Println()
		}
		printing.PrintBlue("|")
		for _, Col := range Row {
			if Col == 0 {
				fmt.Print("O")
			} else if Col == 1 {
				printing.PrintRed("O")
			} else {
				printing.PrintYellow("O")
			}
			printing.PrintBlue("|")
		}
		printing.PrintBlue("\n")
	}
}

// GameOver determines whether game state is terminal
func (c4 *Connect4) GameOver() bool {
	return ((c4.Nump1Pieces == 0 && c4.Nump2Pieces == 0) || c4.Winner != nil)
}

// PrintTurn Prints the current turn of the player
func (c4 *Connect4) PrintTurn() {
	fmt.Println("Player " + strconv.Itoa(c4.Turn) + " Move")
}

// PrintWinner prints the winner
func (c4 *Connect4) PrintWinner() {
	if c4.Winner != nil {
		fmt.Println("Player " + strconv.Itoa(*(c4.Winner)) + " has won the game.")
		return
	}

	fmt.Println("This game has resulted in a draw.")
}

// RandomMove performs a random move in the game
func (c4 *Connect4) RandomMove(print bool) {
	playCol := rand.Intn(7)
	err := c4.Move(playCol+1, print)
	for err != nil {
		err = c4.Move(rand.Intn(7)+1, print)
	}

}

// isWinngMove determines if the last move made was a winning move
func (c4 *Connect4) isWinningMove(player int) {
	var coords []Coordinate
	if player == 1 {
		coords = c4.P1positions
	} else {
		coords = c4.P2positions
	}
	for _, coord := range coords {
		// Connect4 to the right
		if coord.Col+3 < 7 {
			win := true
			for i := coord.Col; i < coord.Col+4; i++ {
				win = win && c4.Layout[coord.Row][i] == player
			}
			if win == true {
				val := player
				c4.Winner = &val
				return
			}
		}
		// Connect4 to the left
		if coord.Col-3 >= 0 {
			win := true
			for i := coord.Col; i > coord.Col-4; i-- {
				win = win && c4.Layout[coord.Row][i] == player
			}
			if win == true {
				val := player
				c4.Winner = &val
				return
			}
		}
		// Connect4 down
		if coord.Row+3 < 6 {
			win := true
			for i := coord.Row; i < coord.Row+4; i++ {
				win = win && c4.Layout[i][coord.Col] == player
			}
			if win == true {
				val := player
				c4.Winner = &val
				return
			}
		}
		// Connect4 up
		if coord.Row-3 >= 0 {
			win := true
			for i := coord.Row; i > coord.Row-4; i-- {
				win = win && c4.Layout[i][coord.Col] == player
			}
			if win == true {
				val := player
				c4.Winner = &val
				return
			}
		}
		// Connect4 diagonally up and to the right
		if coord.Row+3 < 6 && coord.Col-3 >= 0 {
			win := true
			Col := coord.Col
			for i := coord.Row; i < coord.Row+4; i++ {
				win = win && c4.Layout[i][Col] == player
				Col--
			}
			if win == true {
				val := player
				c4.Winner = &val
				return
			}
		}
		// Connect4 diagonally up and to the left
		if coord.Row+3 < 6 && coord.Col+3 < 7 {
			win := true
			Col := coord.Col
			for i := coord.Row; i < coord.Row+4; i++ {
				win = win && c4.Layout[i][Col] == player
				Col++
			}
			if win == true {
				val := player
				c4.Winner = &val
				return
			}
		}
		// Connect4 diagonally down and to the left
		if coord.Row-3 >= 0 && coord.Col+3 < 7 {
			win := true
			Col := coord.Col
			for i := coord.Row; i > coord.Row-4; i-- {
				win = win && c4.Layout[i][Col] == player
				Col++
			}
			if win == true {
				val := player
				c4.Winner = &val
				return
			}
		}
		// Connect4 diagonally down and to the right
		if coord.Row-3 >= 0 && coord.Col-3 >= 0 {
			win := true
			Col := coord.Col
			for i := coord.Row; i > coord.Row-4; i-- {
				win = win && c4.Layout[i][Col] == player
				Col--
			}
			if win == true {
				val := player
				c4.Winner = &val
				return
			}
		}
		if c4.Winner != nil {
			break
		}
	}

}
