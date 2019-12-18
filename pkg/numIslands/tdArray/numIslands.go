package tdArray

import "github.com/solympe/Golang_Training/pkg/numIslands/islandCounter"

// tdArray is struct with input slice
type tdArray struct {
	slice [][]byte
}

// NumIslands is a function that counts the number of islands in a slice
func (t *tdArray) NumIslands(grid [][]byte) int {
	result := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				lenJ := len(grid[i])
				nullNear(grid, i, j, lenJ)
				result++
			}
		}
	}

	return result
}

// GetSlice function returns slice
func (t *tdArray) GetSlice() [][]byte {
	return t.slice
}

// nullNumber is a recursive function that resets cells to 0
func nullNear(grid [][]byte, i int, j int, lenJ int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= lenJ || grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	nullNear(grid, i+1, j, lenJ)
	nullNear(grid, i, j+1, lenJ)
	nullNear(grid, i-1, j, lenJ)
	nullNear(grid, i, j-1, lenJ)
}

// NewtdArray returns new copy of tdArray struct
func NewtdArray(bytes [][]byte) islandCounter.IslandCounter {
	return &tdArray{slice: bytes}
}
