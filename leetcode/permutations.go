// https://leetcode.com/problems/permutations/


func permute(nums []int) [][]int {
	var (
		rs [][]int
		x  []int
	)
	return backtrack(nums, rs, x)
}

func backtrack(nums []int, rs [][]int, x []int) [][]int {
	if len(nums) == 0 {
		rs = append(rs, x)
		return rs
	}
	for i := range nums {
		numsCopy := append([]int{}, nums...)
		xCopy := append([]int{}, x...)
		o := append(numsCopy[:i], numsCopy[i+1:]...)
		xCopy = append(xCopy, nums[i])
		rs = backtrack(o, rs, xCopy)
	}
	return rs
}
