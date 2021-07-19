package main

import "fmt"

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	a:= append([]int{1}, digits...)

	return a
}

func main() {
	a := []int{9, 9, 9, 9}
	fmt.Println(plusOne(a))
}
