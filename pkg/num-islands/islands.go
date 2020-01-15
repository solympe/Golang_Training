package islands

// Islands represents two-dimension array interface
type Islands interface {
	NumIslands() int
	nullNear(grid [][]byte, i int, j int, lenJ int)
}

type islands struct {
	slice [][]byte
}

// NumIslands is a function that counts the number of islands in a slice
func (t *islands) NumIslands() int {
	grid := t.slice
	result := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				lenJ := len(grid[i])
				t.nullNear(grid, i, j, lenJ)
				result++
			}
		}
	}
	return result
}

// nullNear is a recursive function that resets cells to 0
func (t *islands) nullNear(grid [][]byte, i int, j int, lenJ int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= lenJ || grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	t.nullNear(grid, i+1, j, lenJ)
	t.nullNear(grid, i, j+1, lenJ)
	t.nullNear(grid, i-1, j, lenJ)
	t.nullNear(grid, i, j-1, lenJ)
}

// NewIslands returns new copy of islands interface
func NewIslands(bytes [][]byte) Islands {
	return &islands{slice: bytes}
}
