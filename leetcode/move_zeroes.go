// https://leetcode.com/problems/move-zeroes/


func moveZeroes(nums []int)  {
	length := len(nums)
	if length < 2 {
		return
	}
	i, j := 0, 1
	for i < length && j < length {
		if nums[i] == 0 {
			for nums[j] == 0 && j < length-1 {
				j++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		i++
		j++
	}
}
