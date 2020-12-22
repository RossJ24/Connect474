package connect4

// Coordinate a struct to represesnt positions in a Grid
type Coordinate struct {
	Row int
	Col int
}

// newCoord Coordinate constructor
func newCoord(x, y int) Coordinate {
	return Coordinate{x, y}
}
