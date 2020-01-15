package data

import "sort"

// Data ...
type Data interface {
	Merge(data1 Data, data2 Data) []int
	getNums() []int
	getVal() int
}

type data struct {
	nums []int
	val  int
}

// Merge ...
func (d *data) Merge(data1 Data, data2 Data) []int {
	nums1 := data1.getNums()
	nums2 := data2.getNums()
	m := data1.getVal()
	n := data2.getVal()

	if n == 0 {
		return nums1
	}
	i := 0
	j := 0
	for i != m {
		if nums1[i] > nums2[j] {
			tmp := nums2[j]
			nums2[j] = nums1[i]
			nums1[i] = tmp
			sort.Ints(nums2)
		}
		i++
	}
	for i := 0; i < n; i++ {
		nums1[m] = nums2[i]
		m++
	}
	return nums1
}

func (d *data) getNums() []int {
	return d.nums
}

func (d *data) getVal() int {
	return d.val
}

// NewData ...
func NewData(nums []int, val int) Data {
	return &data{
		nums: nums,
		val:  val,
	}
}
