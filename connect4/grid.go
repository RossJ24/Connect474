package connect4

// Grid  is an array of integers representing a Connect4 grid
type Grid [6][7]int

// CopyGrid returns a deep copy of a Grid
func (grid *Grid) CopyGrid() Grid {
	dup := new(Grid)
	for i := range grid {
		copy(dup[i][:], grid[i][:])
	}
	return *dup
}
