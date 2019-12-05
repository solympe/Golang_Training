package main

import "fmt"

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

func main() {
	var i = []int{1, 2, 3, 4}
	fmt.Println(twoSum(i, 4))
}