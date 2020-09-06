// https://leetcode.com/problems/missing-number/
package main

// sum method
//func missingNumber(nums []int) int {
//	sum := 0
//	for i := 0; i < len(nums); i++ {
//		sum += nums[i]
//	}
//	noMissingSum := 0
//	for i := 1; i < len(nums)+1; i++ {
//		noMissingSum += i
//	}
//	return noMissingSum - sum
//}

// bit op method
func missingNumber(nums []int) int {
	rv := len(nums)
	for i := 0; i < len(nums); i++ {
		rv ^= i ^ nums[i]
	}
	return rv
}
