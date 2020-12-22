package tree

import (
	"testing"

	"github.com/RossJ24/connect474/connect4"
)

// TestTree ensures that the Tree is being created properly
func TestTree(t *testing.T) {
	c4 := connect4.NewConnect4()
	tree := NewTree(&c4)
	if tree.Root.Parent != nil {
		t.Errorf("Root is not nil")
	}
}

// TestNode ensures that the basic functions of the Node are working properly
func TestNode(t *testing.T) {
	c4 := connect4.NewConnect4()
	n := NewNode(nil, 5, &c4)
	for !n.FullyExpanded() {
		uelenb4 := len(n.Unexpanded)
		kid := n.Expand()
		uelenafter := len(n.Unexpanded)
		if uelenafter != uelenb4-1 {
			t.Errorf("Unexpanded length is incorrect")
		}
		if kid == nil {
			t.Errorf("Child node is nil")
		}
	}

}
