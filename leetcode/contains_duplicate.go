// https://leetcode.com/problems/contains-duplicate/


func containsDuplicate(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return false
	}
	c := make(map[int]int, length)
	for i := 0; i < length; i++ {
		if c[nums[i]] == 1 {
			return true
		}
		c[nums[i]]++
	}
	return false
}
