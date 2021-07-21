package main

import "fmt"

func firstUniqChar(s string) int {
	var test [26]int
	for i := 0; i < len(s); i++ {
		test[s[i] - 'a'] ++
	}
	for i := 0; i < len(s); i++ {
		if test[s[i] - 'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
