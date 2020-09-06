// https://leetcode.com/problems/search-in-rotated-sorted-array/
package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[start] <= nums[mid] {
			if target < nums[mid] && target >= nums[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
