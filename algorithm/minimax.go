package algorithm

import (
	"math"

	"github.com/RossJ24/connect474/connect4"
)

//MiniMax algorithm for Connect 4
func MiniMax(c4 *connect4.Connect4) int {
	connect4 := c4.CopyConnect4()

	possibleMoves := connect4.PossibleMoves()
	if len(possibleMoves) == 0 || connect4.Winner != nil {
		return 1
	}

	_, action := MiniMaxAB(connect4, 0, math.Inf(-1), math.Inf(1))

	return action
}

//MiniMaxAB is minimax with Alpha Beta
func MiniMaxAB(c4 *connect4.Connect4, depth int, alpha float64, beta float64) (float64, int) {
	possibleMoves := c4.PossibleMoves()
	if len(possibleMoves) == 0 || c4.Winner != nil {
		return (float64)(c4.GetReward()) * 1000, 1
	}
	if depth == 5 {
		hval, d := (float64)(heuristic(c4)), 1
		return hval, d
	}
	if c4.Turn == 1 {
		best, bestdex := math.Inf(-1), 0
		for _, action := range possibleMoves {
			c4.Move(action, false)
			val, _ := MiniMaxAB(c4, depth+1, alpha, beta)

			if val > best {

				best = val
				bestdex = action
			}
			if best > alpha {
				alpha = best
			}
			c4.UndoMove()
			if beta <= alpha {

				break
			}

		}
		return best, bestdex
	}

	best, bestdex := math.Inf(1), 0
	for _, action := range possibleMoves {
		c4.Move(action, false)
		val, _ := MiniMaxAB(c4, depth+1, alpha, beta)
		if val < best {
			best = val
			bestdex = action
		}
		if best < alpha {
			beta = best
		}
		c4.UndoMove()
		if beta <= alpha {

			break
		}

	}

	return best, bestdex
}

// heuristic is the heuristic for determining the value of non-terminal states
func heuristic(c4 *connect4.Connect4) int {
	player1value := heuristichelper(3, 1, c4)*2 + heuristichelper(2, 1, c4)
	player2value := heuristichelper(3, 2, c4)*2 + heuristichelper(2, 2, c4)
	if c4.Turn == 2 {
		return player1value - player2value
	}
	return player2value - player1value
}

// heuristichelper is the function that actually calculates the value for the heuristic function
func heuristichelper(inaRow int, player int, c4 *connect4.Connect4) int {
	var coords []connect4.Coordinate
	num := 0
	if player == 1 {
		coords = c4.P1positions
	} else {
		coords = c4.P2positions
	}
	for _, coord := range coords {
		if coord.Col+inaRow-1 < 7 {
			connect := true
			for i := coord.Col; i < coord.Col+inaRow; i++ {
				connect = connect && c4.Layout[coord.Row][i] == player
			}
			if connect == true {
				num++
			}
		}
		if coord.Col-inaRow-1 >= 0 {
			connect := true
			for i := coord.Col; i > coord.Col-inaRow; i-- {
				connect = connect && c4.Layout[coord.Row][i] == player
			}
			if connect == true {
				num++
			}
		}
		if coord.Row+inaRow-1 < 6 {
			connect := true
			for i := coord.Row; i < coord.Row+inaRow; i++ {
				connect = connect && c4.Layout[i][coord.Col] == player
			}
			if connect == true {
				num++
			}
		}
		if coord.Row-inaRow-1 >= 0 {
			connect := true
			for i := coord.Row; i > coord.Row-inaRow; i-- {
				connect = connect && c4.Layout[i][coord.Col] == player
			}
			if connect == true {
				num++
			}
		}
		if coord.Row+inaRow-1 < 6 && coord.Col-inaRow-1 >= 0 {
			connect := true
			Col := coord.Col
			for i := coord.Row; i < coord.Row+inaRow; i++ {
				connect = connect && c4.Layout[i][Col] == player
				Col--
			}
			if connect == true {
				num++
			}
		}
		if coord.Row+inaRow-1 < 6 && coord.Col+inaRow-1 < 7 {
			connect := true
			Col := coord.Col
			for i := coord.Row; i < coord.Row+inaRow; i++ {
				connect = connect && c4.Layout[i][Col] == player
				Col++
			}
			if connect == true {
				num++
			}
		}

		if coord.Row-inaRow-1 >= 0 && coord.Col+inaRow-1 < 7 {
			connect := true
			Col := coord.Col
			for i := coord.Row; i > coord.Row-inaRow; i-- {
				connect = connect && c4.Layout[i][Col] == player
				Col++
			}
			if connect == true {
				num++
			}
		}
		if coord.Row-inaRow-1 >= 0 && coord.Col-inaRow-1 >= 0 {
			connect := true
			Col := coord.Col
			for i := coord.Row; i > coord.Row-inaRow; i-- {
				connect = connect && c4.Layout[i][Col] == player
				Col--
			}
			if connect == true {
				num++
			}
		}

	}
	return num
}
