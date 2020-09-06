// https://leetcode.com/problems/maximum-subarray/
package main

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	maxValue, curValue := nums[0], nums[0]
	for i := 1; i < length; i++ {
		sum := curValue + nums[i]
		if sum > nums[i] {
			curValue = sum
		} else {
			curValue = nums[i]
		}
		if curValue > maxValue {
			maxValue = curValue
		}
	}
	return maxValue
}
