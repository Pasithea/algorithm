// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package main

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	j := 0
	for i := 1; i < length; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
