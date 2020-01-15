package nums

// Nums represents nums Interface
type Nums interface {
	TwoSum() []int
}

type nums struct {
	nums   []int
	target int
}

// TwoSum ...
func (g *nums) TwoSum() []int {
	var result []int

	for i := 0; i < len(g.nums)-1; i++ {
		for j := i + 1; j < len(g.nums); j++ {
			if g.nums[i]+g.nums[j] == g.target {
				result = append(result, i, j)
				return result
			}
		}
	}
	return result
}

// NewNums returns new nums interface
func NewNums(numbs []int, target int) Nums {
	return &nums{nums: numbs, target: target}
}
