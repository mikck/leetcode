package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	length := len(s)
	i := 0
	j := length - 1
	for i <= j {
		for !isalnum(s[i]) && i<j {
			i++
		}
		for !isalnum(s[j]) && i<j {
			j--
		}

		if s[i] != s[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}
func isalnum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}