// https://leetcode.com/problems/house-robber/


var m []int

func rob(nums []int) int {
	length := len(nums)
	m = make([]int, length)
	for i := 0; i < length; i++ {
		m[i] = -1
	}
	return add(nums, 0, length)
}

func add(nums []int, start, length int) int {
	if start >= length {
		return 0
	}
	if m[start] >= 0 {
		return m[start]
	}
	sum1 := nums[start] + add(nums, start+2, length)
	sum2 := add(nums, start+1, length)
	var v int
	if sum1 > sum2 {
		v = sum1
	} else {
		v = sum2
	}
	m[start] = v
	return v
}


// dp solution.
//
//func rob(nums []int) int {
//	length := len(nums)
//	if length == 0 {
//		return 0
//	}
//	if length == 1 {
//		return nums[0]
//	}
//	max := func (a, b int) int {
//		if a > b {
//			return a
//		}
//		return b
//	}
//	m := make([]int, length)
//	m[0] = nums[0]
//	m[1] = max(nums[0], nums[1])
//	for i := 2; i < length; i++ {
//		m[i] = max(m[i-1], m[i-2]+nums[i])
//	}
//	return m[length-1]
//}
