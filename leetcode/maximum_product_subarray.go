// https://leetcode.com/problems/maximum-product-subarray/


func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	rv, maxValue, minValue := nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			minValue, maxValue = maxValue, minValue
		}
		maxValue = max(maxValue * nums[i], nums[i])
		minValue = min(minValue * nums[i], nums[i])
		rv = max(maxValue, rv)
	}
	return rv
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
