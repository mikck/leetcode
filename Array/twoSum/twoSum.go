package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := res[nums[i]]; ok {
			return []int{res[nums[i]], i}
		} else {
			res[target-nums[i]] = i
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3,3}, 6))
}
