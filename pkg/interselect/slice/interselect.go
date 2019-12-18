package slice

import ss "github.com/solympe/Golang_Training/cmd/interselect/SliceSolver"

type slice struct {
	data []int
}

// Interselect compute intersection of 2 slices
func Intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	for _, num := range nums1 {
		for j, num2 := range nums2 {
			if num == num2 {
				result = append(result, num)
				nums2[j] = -num2
				break
			}
		}
	}
	return result
}

// GetSlice returns data from slice struct
func (s *slice) GetSlice() []int {
	return s.data
}

// NewSlice returns new copy of slice
func NewSlice(sliceIn []int) ss.SliceSolver {
	return &slice{data: sliceIn}
}
