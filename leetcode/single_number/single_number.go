// https://leetcode.com/problems/single-number/
package main

func singleNumber(nums []int) int {
	rv := nums[0]
	for i := 1; i < len(nums); i++ {
		rv ^= nums[i]
	}
	return rv
}
