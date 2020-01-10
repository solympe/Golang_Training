<b>TwoSum from leetcode</b>

Version for leetcode: 

<a href ="https://leetcode.com/problems/two-sum/">LINK</a>

```
func twoSum(nums []int, target int) []int {

	var result []int

	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				result = append(result, i, j)
				return result
			}
		}
	}
	return result
}
```

Given an array of integers, return indices of the two numbers such that they add up to a specific target.<br>
You may assume that each input would have exactly one solution, and you may not use the same element twice.