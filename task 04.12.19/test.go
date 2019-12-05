package main

import "fmt"


func twoSum(nums []int, target int) []int {

	var result = []int{}

	for i := 0; i < len(nums) - 1; i++{
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

//Given nums = [2, 7, 11, 15], target = 9

arr := []int{0, 2, 3, 6, 2, 2, 2, 7, 11, 15, 9}

fmt.Println(twoSum(arr, 9))

}

