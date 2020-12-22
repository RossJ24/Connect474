package algorithm

import (
	"testing"

	"github.com/RossJ24/connect474/connect4"
)

// TestMCTS tests Monte Carlo Tree Search agent
func TestMCTS(t *testing.T) {
	wins := 0
	for i := 0; i < 1000; i++ {
		flip := true
		c4 := connect4.NewConnect4()
		for !c4.GameOver() {
			if flip {
				c4.RandomMove(false)
			} else {
				c4.Move(MCTS(&c4), false)
			}
			flip = !flip
		}

		if c4.Winner != nil && *c4.Winner == 2 {
			wins++
		}
	}
	percentage := ((float64)(wins) / (float64)(1000))
	if percentage < .9 {
		t.Errorf("Win percentage was too low")
	}
	if percentage > 1 {
		t.Errorf("Win percentage was too high")
	}

}

// TestMMAB tests MiniMax agent with Alpha Beta pruning
func TestMMAB(t *testing.T) {
	wins := 0
	for i := 0; i < 1000; i++ {
		flip := true
		c4 := connect4.NewConnect4()
		for !c4.GameOver() {
			if flip {
				c4.RandomMove(false)
			} else {
				c4.Move(MiniMax(&c4), false)
			}
			flip = !flip

		}
		if c4.Winner != nil && *c4.Winner == 2 {
			wins++
		}
	}
	percentage := ((float64)(wins) / (float64)(1000))
	if percentage < .9 {
		t.Errorf("Win percentage was too low")
	}
	if percentage > 1 {
		t.Errorf("Win percentage was too high")
	}
}
