package main

import "fmt"

func main() {
	nums := []int{4, 3, 7, 8, 2, 5, 0, 1, 9, 6}
	sort(nums)
	fmt.Println(nums)
}

func sort(nums []int) {
	if len(nums) < 2 {
		return
	}
	qsort(nums, 0, len(nums)-1)
}

func qsort(nums []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(nums, start, end)
	qsort(nums, start, i-1)
	qsort(nums, i+1, end)
}

func partition(nums []int, start, end int) int {
	pivot := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}
