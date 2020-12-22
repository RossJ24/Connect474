package tree

import "github.com/RossJ24/connect474/connect4"

// Node : the nodes of the tree
type Node struct {
	// Pointer to parent node
	Parent *Node
	// Depth of the node in the tree
	Depth int
	// The connect4 game being represented in the data
	Data *connect4.Connect4
	// The number of times the node has been visited
	Visits int
	// The total reward for the node
	Totalreward int
	// The children of the node
	Children []*Node
	// The unexpanded children of the node
	Unexpanded []int
	// The player at that node
	Player int
}

// NewNode makes a new node
func NewNode(parent *Node, depth int, connect4 *connect4.Connect4) *Node {
	return &Node{Parent: parent,
		Depth:       depth,
		Data:        connect4.CopyConnect4(),
		Visits:      0,
		Totalreward: 0,
		Children:    make([]*Node, 7),
		Unexpanded:  connect4.PossibleMoves(),
		Player:      connect4.Turn,
	}
}

// Expand expands the node's children
func (n *Node) Expand() *Node {
	action := n.Unexpanded[0]
	n.Unexpanded = n.Unexpanded[1:]
	c4 := n.Data.CopyConnect4()
	c4.Move(action, false)
	newn := NewNode(n, n.Depth+1, c4)
	n.Children[action-1] = newn
	return newn
}

// Visit increments the amount of visits a node has
func (n *Node) Visit() {
	n.Visits++
}

// AddReward increments the total reward
func (n *Node) AddReward(reward int) {
	n.Totalreward += reward
}

// AvgReward returns the average reward of the node
func (n *Node) AvgReward() float64 {
	return (float64)(n.Totalreward) / (float64)(n.Visits)
}

// FullyExpanded determines whether the node has all of it's children
func (n *Node) FullyExpanded() bool {
	return len(n.Unexpanded) == 0
}
