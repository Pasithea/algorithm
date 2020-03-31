// https://leetcode.com/problems/increasing-triplet-subsequence/


func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	maxInt := (1 << 63) - 1
	n1, n2 := maxInt, maxInt
	for i := 0; i < len(nums); i++ {
		if nums[i] <= n1 {
			n1 = nums[i]
		} else if nums[i] <= n2 {
			n2 = nums[i]
		} else {
			return true
		}
	}
	return false
}
