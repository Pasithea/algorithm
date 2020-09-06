// https://leetcode.com/problems/3sum/
package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	rv := [][]int{}
	length := len(nums)
	if length < 3 {
		return rv
	}
	sort.Ints(nums)
	for i := 0; i < length-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		lo, hi := i+1, length-1
		for lo < hi {
			if -nums[i] == nums[lo]+nums[hi] {
				rv = append(rv, []int{nums[i], nums[lo], nums[hi]})
				lo++
				hi--
				for lo < hi && nums[lo] == nums[lo-1] {
					lo++
				}
				for lo < hi && nums[hi] == nums[hi+1] {
					hi--
				}
			} else if -nums[i] < nums[lo]+nums[hi] {
				hi--
			} else {
				lo++
			}
		}
	}
	return rv
}
