package tree

import (
	"github.com/RossJ24/connect474/connect4"
)

// Tree represents the game tree
type Tree struct {
	Root *Node
}

// NewTree makes a new Tree
func NewTree(connect4 *connect4.Connect4) *Tree {
	root := NewNode(nil, 0, connect4)
	return &Tree{Root: root}
}
