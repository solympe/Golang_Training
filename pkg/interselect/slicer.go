package interselect

// Slicer represents an interface for slice struct
type Slicer interface {
	GetSlice() []int
	Intersect(nums1 []int, nums2 []int) []int
}
