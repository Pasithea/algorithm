// https://leetcode.com/problems/longest-increasing-subsequence/


func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				count[i] = max(count[i], count[j]+1)
			}
		}
	}
	m := 0
	for _, v := range count {
		if v > m {
			m = v
		}
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
