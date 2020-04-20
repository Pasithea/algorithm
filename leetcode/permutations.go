// https://leetcode.com/problems/permutations/


func permute(nums []int) [][]int {
	var (
		rs [][]int
		x  []int
	)
	return backtrack(nums, x, rs)
}

func backtrack(nums, x []int, rs [][]int) [][]int {
	if len(nums) == 0 {
		rs = append(rs, x)
		return rs
	}
	for i := range nums {
		nums[i], nums[0] = nums[0], nums[i]
		xCopy := append([]int{}, x...)
		xCopy = append(xCopy, nums[0])
		rs = backtrack(nums[1:], xCopy, rs)
		nums[i], nums[0] = nums[0], nums[i]
	}
	return rs
}
