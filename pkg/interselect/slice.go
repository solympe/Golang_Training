package interselect

type slice struct {
	data []int
}

// GetSlice returns data from slice struct
func (s *slice) GetSlice() []int {
	return s.data
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

// NewSlicer returns new copy of Slicer interface
func NewSlicer(sliceIn []int) Slicer {
	return &slice{data: sliceIn}
}
