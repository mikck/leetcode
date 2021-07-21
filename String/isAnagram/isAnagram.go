package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	lengthS := len(s)
	lengthT := len(t)
	if lengthS != lengthT {
		return false
	}
	res := make(map[byte]int)
	for i := 0; i < lengthS; i++ {
		res[s[i]]++
	}
	for i := 0; i < lengthT; i++ {
		if _, ok := res[t[i]]; !ok {
			return false
		} else {
			res[t[i]]--
		}
	}
	for _, v := range res {
		if v > 0 {
			return false
		}
	}
	//哈希表
	//var c1, c2 [26]int
	//for _, ch := range s {
	//	c1[ch-'a']++
	//}
	//for _, ch := range t {
	//	c2[ch-'a']++
	//}
	//return c1 == c2

	//排序
	//s1, s2 := []byte(s), []byte(t)
	//sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	//sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	//return string(s1) == string(s2)

	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
