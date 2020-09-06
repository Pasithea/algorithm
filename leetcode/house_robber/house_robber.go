// https://leetcode.com/problems/house-robber/
package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	count := make([]int, len(nums)+1)
	count[1] = nums[0]
	for i := 2; i < len(nums)+1; i++ {
		count[i] = max(count[i-1], count[i-2]+nums[i-1])
	}
	return count[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
