// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	start, end := 0, len(nums)-1
	i, j := -1, -1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			if (mid-1) >= start && nums[mid-1] == target {
				end = mid - 1
			} else {
				i = mid
				break
			}
		}
	}
	if i != -1 {
		start, end = i, len(nums)-1
		for start <= end {
			mid := (start + end) / 2
			if nums[mid] > target {
				end = mid - 1
			} else if nums[mid] < target {
				start = mid + 1
			} else {
				if (mid+1) <= end && nums[mid+1] == target {
					start = mid + 1
				} else {
					j = mid
					break
				}
			}
		}
		if j == -1 {
			j = 0
		}
	}
	return []int{i, j}
}
