package main

import (
	"sort"
)

func Merge(nums1 []int, m int, nums2 []int, n int) []int{
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
