package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	length := len(nums)
	if nums != nil && length != 0 {
		index := 0
		for i := 0; i < length; i++ {
			if nums[i] != 0 {
				nums[index] = nums[i]
				index++
			}
		}
		for index < length {
			nums[index] = 0
			index++
		}
	}

	fmt.Println(nums)
}


func main() {
	a := []int{0, 9, 0, 9}
	moveZeroes(a)
}
