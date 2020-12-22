package algorithm

import (
	"math"

	"github.com/RossJ24/connect474/connect4"
	"github.com/RossJ24/connect474/tree"
)

// MCTS (Monte Carlo Tree Search for Connect 4)
func MCTS(c4 *connect4.Connect4) int {
	t := tree.NewTree(c4.CopyConnect4())
	for i := 0; i < 1000; i++ {
		n := traverse(t)
		if n.Data.GameOver() {
			reward := n.Data.GetWinner()
			Backpropagate(t, n, reward)
			continue
		}
		reward := Simulate(n)
		Backpropagate(t, n, reward)
	}
	return BestAction(t) + 1
}

// UCB function for the Node
func UCB(n *tree.Node) float64 {
	rj := n.AvgReward()
	nj := (float64)(n.Visits)
	t := (float64)(n.Parent.Visits)
	exploration := math.Sqrt((5.0 * math.Log(t)) / nj)
	if n.Parent.Player == 1 {
		return rj + exploration
	}
	return rj - exploration
}

// NextNode returns the index of the child node to choose
func NextNode(n *tree.Node) int {
	bestdex, best := 0, 0.0

	if n.Player == 1 {
		best = math.Inf(-1)
	} else {
		best = math.Inf(1)
	}
	for i, node := range n.Children {
		if node == nil {
			continue
		}
		if n.Player == 1 {
			val := UCB(node)
			if val > best {
				best = val
				bestdex = i
			}
		} else {
			val := UCB(node)
			if val < best {
				best = val
				bestdex = i
			}
		}
	}
	return bestdex
}

// Simulate a playout to a terminal state
func Simulate(n *tree.Node) int {
	c4 := n.Data.CopyConnect4()

	for !c4.GameOver() {
		c4.RandomMove(false)
	}
	if c4.Winner == nil {
		return 0
	} else if *(c4.Winner) == 1 {
		return 1
	} else if *(c4.Winner) == 2 {
		return -1
	} else {
		return 0
	}

}

// Backpropagate results to the root of the tree
func Backpropagate(t *tree.Tree, n *tree.Node, r int) {
	curr := n
	for curr != nil {
		curr.Visit()
		curr.AddReward(r)
		curr = curr.Parent
	}
}

// BestAction gives the best action that a player could make
func BestAction(t *tree.Tree) int {
	best, bestdex := 0.0, 0
	if t.Root.Player == 1 {
		best = math.Inf(-1)
	} else {
		best = math.Inf(1)
	}
	for i, node := range t.Root.Children {
		if node == nil {
			continue
		}
		if t.Root.Player == 1 {
			if node.AvgReward() > best {
				best = node.AvgReward()
				bestdex = i
			}
		} else {
			if node.AvgReward() < best {
				best = node.AvgReward()
				bestdex = i
			}
		}
	}
	return bestdex
}

//traverse the tree until there's an unexpanded node
func traverse(t *tree.Tree) *tree.Node {
	curr := t.Root
	for curr.FullyExpanded() && !curr.Data.GameOver() {
		curr = curr.Children[NextNode(curr)]
	}
	if curr.Data.GameOver() {
		return curr
	}

	return curr.Expand()

}
