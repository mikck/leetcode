package main

import "fmt"

func strStr(haystack string, needle string) int {
	lengthH := len(haystack)
	lengthN := len(needle)
	i := 0
	j := 0
	for i < lengthH && j < lengthN {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j == lengthN {
		return i - j
	}

	return -1
}

func main() {
	fmt.Println(strStr("", ""))
}
