// https://leetcode.com/problems/rotate-array/
package main

//func rotate(nums []int, k int)  {
//	length := len(nums)
//	k = k % length
//	if length < 2 || k == 0 {
//		return
//	}
//	for i := 0; i < k; i++ {
//		tmp := nums[length-1]
//		for j := length-1; j > 0; j-- {
//			nums[j] = nums[j-1]
//		}
//		nums[0] = tmp
//	}
//}

func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	if length < 2 || k == 0 {
		return
	}
	reverse(nums, 0, length-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, length-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
