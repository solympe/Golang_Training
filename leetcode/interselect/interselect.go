package interselect

type SliceSolver interface {
	Intersect(nums1 []int, nums2 []int) []int
	GetSlice() []int
}

type slice struct {
	data []int
}

func (s *slice)Intersect(nums1 []int, nums2 []int) []int {
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

func (s *slice)GetSlice() []int{
	return s.data
}

func NewSlice (sliceIn []int) SliceSolver{
	return &slice{data: sliceIn}
}