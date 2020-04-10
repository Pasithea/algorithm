// https://leetcode.com/problems/sort-colors/


func sortColors(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	head, cur, tail := 0, 0, len(nums) - 1
	for cur <= tail {
		if nums[cur] == 0 {
			nums[head], nums[cur] = nums[cur], nums[head]
			head++
			cur++
		} else if nums[cur] == 1 {
			cur++
		} else if nums[cur] == 2 {
			nums[tail], nums[cur] = nums[cur], nums[tail]
			tail--
		}
	}
}
