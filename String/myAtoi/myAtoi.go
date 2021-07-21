package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	length := len(s)
	index := 0
	// 先去除空格
	for index < length && s[index] == ' ' {
		index++
	}
	// 极端情况 "  " 和""
	if index >= length {
		return 0
	}
	// 再判断符号
	sign := 1
	if s[index] == '-' || s[index] == '+' {
		if s[index] == '-' {
			sign = -1
		}
		index++
	}
	result := 0
	for index < length {
		num := s[index] - '0'
		if num > 9 || num < 0 {
			break
		}
		result = result*10 + int(num)
		// 越界后，数值和期望数值发生变化，取余再除10获取原始值，比对判断
		if result>2147483647 {
			if sign > 0 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		index++
	}
	return result * sign
}

func main() {
	fmt.Println(myAtoi("-2147483647"))
}
