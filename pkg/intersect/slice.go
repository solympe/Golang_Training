package slice

// Slice represents an interface for slice struct
type Slice interface {
	GetSlice() []int
	Intersect(nums1 []int, nums2 []int) []int
}

type slice struct {
	data []int
}

// GetSlice returns data from slice struct
func (s *slice) GetSlice() []int {
	return s.data
}

// Intersect compute intersection of 2 slices
func (s *slice) Intersect(nums1 []int, nums2 []int) []int {
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

// NewSlice returns new copy of Slice
func NewSlice(sliceIn []int) Slice {
	return &slice{data: sliceIn}
}
