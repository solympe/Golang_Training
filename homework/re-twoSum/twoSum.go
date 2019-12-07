package main

import "fmt"

//Given nums = [2, 7, 11, 15], target = 9,
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1]

func main() {

	input := []int{2, 2, 7, 11, 15}
	target := 9
	fmt.Println("Output:",twoSum(input, target))

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		c := target - num
		check, ok := m[c]
		if ok == true {
			return []int{check, i}
		}
		m[num] = i
	}

	return []int{}
}