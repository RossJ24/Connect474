package connect4

import (
	"testing"
)

func TestConnect4(t *testing.T) {
	c4 := NewConnect4()
	c4.Move(5, false)
	c42 := c4.CopyConnect4()
	if c4.Turn != c42.Turn {
		t.Errorf("Turns are not being copied correctly")
	}
	c42.Move(7, false)
	if c42.lastPlayerTurn() != c4.Turn {
		t.Errorf("Last Player Turn has an Error")
	}
	if len(c4.P1positions) != len(c42.P1positions) {
		t.Errorf("Player positions are not being copied correctly.")
	}
	if c42.GetReward() != 0 {
		t.Errorf("Reward is not 0 in a non-terminal state")
	}
	c42.UndoMove()
	if len(c42.P2positions) != 0 {
		t.Errorf("Player2's move was not undone")
	}
}

func TestGrid(t *testing.T) {
	var grid Grid
	grid[1][1] = 1
	grid2 := grid.CopyGrid()
	if grid2[1][1] != 1 {
		t.Errorf("Incorrect copy")
	}
	grid[5][5] = 1
	if grid[5][5] == grid2[5][5] {
		t.Errorf("Same reference")
	}
}
