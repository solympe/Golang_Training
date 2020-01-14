package islands

// IslandsCounter represents two-dimension array interface
type IslandsCounter interface {
	NumIslands(grid [][]byte) int
	GetSlice() [][]byte
}
