// https://leetcode.com/problems/intersection-of-two-arrays-ii/
package main

import (
	"sort"
)

// use map
//func intersect(nums1 []int, nums2 []int) []int {
//	m := make(map[int]int)
//	rv := []int{}
//	for _, v := range nums1 {
//		m[v]++
//	}
//	for _, v := range nums2 {
//		if m[v] > 0 {
//			rv = append(rv, v)
//			m[v]--
//		}
//	}
//	return rv
//}

// use sort
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	rv := []int{}
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			rv = append(rv, nums1[i])
			i++
			j++
		}
	}
	return rv
}
