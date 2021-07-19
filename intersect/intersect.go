package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var c []int
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			// 如果i指向的值大于j指向的值，说明j指向的值
			// 小了，j往后移一步
			j++
		} else {
			// 如果i和j指向的值相同，说明这两个值是重复的，
			// 把他加入到集合list中，然后i和j同时都往后移一步
			c = append(c, nums1[i])
			i++
			j++
		}
	}
	return c
}

func main() {
	a := []int{7, 5, 5, 6}
	b := []int{0, 9, 4, 5, 6}
	c := intersect(a, b)
	fmt.Println(c)
}
