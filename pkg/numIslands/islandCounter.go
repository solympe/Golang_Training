package islands

// IslandCounter represents two-dimension array interface
type IslandCounter interface {
	NumIslands(grid [][]byte) int
	GetSlice() [][]byte
}
