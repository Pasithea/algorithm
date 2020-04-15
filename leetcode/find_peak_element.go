// https://leetcode.com/problems/find-peak-element/


func findPeakElement(nums []int) int {
	i, j := 0, len(nums) - 1
	for i <= j {
		if i == j {
			return j
		}
		mid := (i + j) / 2
		if mid < j && nums[mid] < nums[mid+1] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return j
}
