package modes

import (
	"fmt"

	"github.com/RossJ24/connect474/algorithm"
	"github.com/RossJ24/connect474/connect4"
)

// errcheck checks if there is an error
func errcheck(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Help gives information about how to run the program
func Help() {
	fmt.Println("Connect474")
	fmt.Println("Usage: ./connect474 --{Mode} {Algorithm} {# of games}")
	fmt.Println("Mode : \n\t--R : Watch two random agents play \n\t--PvC : Person vs Computer \n\t--CvC : Computer vs Computer")
	fmt.Println("Algorithm : \n\tMCTS :Monte Carlo Tree Search\n\tMMAB : Minimax with Alpha Beta")
	fmt.Println("# of games : int of number of games to run")
	fmt.Println("example:  ./connect474 --CvC MCTS 100")
}

// RandomPlay nmakes you watch two random agents play eachother
func RandomPlay() {
	c4 := connect4.NewConnect4()
	c4.Print()
	for !c4.GameOver() {
		c4.RandomMove(true)
	}
}

// PvP Player vs Player mode
func PvP() {
	c4 := connect4.NewConnect4()
	c4.Print()
	for !c4.GameOver() {
		c4.PrintTurn()
		var i int
		_, err := fmt.Scanf("%d", &i)
		if errcheck(err) {
			err = c4.Move(i, true)
			errcheck(err)
		}
	}
	c4.PrintWinner()
}

// PvC Player vs Computer Mode
func PvC(algo bool) {
	c4 := connect4.NewConnect4()
	if algo {
		c4.Print()
		for !c4.GameOver() {
			c4.PrintTurn()
			err := c4.Move(algorithm.MCTS(&c4), true)
			if c4.GameOver() {
				break
			}
			c4.PrintTurn()

			var i int
			_, err = fmt.Scanf("%d", &i)
			if errcheck(err) {
				err = c4.Move(i, true)
				errcheck(err)
			}

		}
		c4.PrintWinner()
	} else {
		c4.Print()
		for !c4.GameOver() {
			c4.PrintTurn()
			c4.Move(algorithm.MiniMax(&c4), true)
			if c4.GameOver() {
				break
			}
			c4.PrintTurn()

			var i int
			_, err := fmt.Scanf("%d", &i)
			if errcheck(err) {
				err = c4.Move(i, true)
				errcheck(err)
			}

		}
		c4.PrintWinner()
	}
}

// CvC Computer vs Computer Mode
func CvC(algo bool, games int) {
	wins := 0
	if algo {
		for i := 0; i < games; i++ {
			flip := true
			c4 := connect4.NewConnect4()
			for !c4.GameOver() {
				if flip {
					c4.RandomMove(false)
				} else {
					c4.Move(algorithm.MCTS(&c4), false)
				}
				flip = !flip
			}
			if c4.Winner != nil && *c4.Winner == 2 {
				wins++
			}
		}
	} else {
		for i := 0; i < games; i++ {
			c4 := connect4.NewConnect4()
			flip := true
			for !c4.GameOver() {
				if flip {
					c4.RandomMove(false)
				} else {
					c4.Move(algorithm.MiniMax(&c4), false)
				}
				flip = !flip
			}
			if c4.Winner != nil && *c4.Winner == 2 {
				wins++
			}
		}
	}
	percentage := ((float64)(wins) / (float64)(games))
	fmt.Print(wins)
	fmt.Print(" ")
	fmt.Println(games)
	fmt.Println(percentage)
}

// CVCAlgos MCTS vs MMAB
func CVCAlgos(games int) {
	wins1, wins2, draws := 0, 0, 0
	for i := 0; i < games; i++ {
		c4 := connect4.NewConnect4()
		flip := true
		for !c4.GameOver() {
			if flip {
				c4.Move(algorithm.MCTS(&c4), false)
			} else {
				c4.Move(algorithm.MiniMax(&c4), false)
			}
			flip = !flip
		}
		if c4.Winner != nil && *c4.Winner == 1 {
			wins1++
		} else if c4.Winner != nil && *c4.Winner == 2 {
			wins2++
		} else {
			draws++
		}
	}
	percentage := ((float64)(wins1) / (float64)(games))
	fmt.Print(wins1)
	fmt.Print(" ")
	fmt.Print(wins2)
	fmt.Print(" ")
	fmt.Print(draws)
	fmt.Print(" ")
	fmt.Println(games)
	fmt.Println(percentage)
}
