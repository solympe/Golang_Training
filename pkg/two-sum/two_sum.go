package two_sum

// TwoSumSolver represents solverInterface
type TwoSumSolver interface {
	TwoSum() []int
}

type givenNums struct {
	nums   []int
	target int
}

// TwoSum ...
func (g *givenNums) TwoSum() []int {
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

// NewGivenNums returns new instance of given nums
func NewGivenNums(nums []int, target int) TwoSumSolver {
	return &givenNums{nums: nums, target: target}
}
