package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var s int = 0
	for x != 0 {
		if s < math.MinInt32/10 || s > math.MaxInt32/10 {
			return 0
		}
		t := x % 10
		newS := s*10 + t
		s=newS
		x/=10
	}
	return s
}

func main() {
	fmt.Println(reverse(1534236469))
}
