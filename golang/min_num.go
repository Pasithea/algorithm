package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{3, 32, 321}
	fmt.Println(minNum(nums))
}

func minNum(nums []int) int {
	sort(nums)
	x := ""
	for _, i := range nums {
		x += strconv.Itoa(i)
	}
	rs, _ := strconv.Atoi(x)
	if x == "" {
		return 0
	}
	return rs
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
	pviot := nums[end]
	i := start
	for j := i; j < end; j++ {
		if less(nums[j], pviot) {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}

func less(a, b int) bool {
	af := strconv.Itoa(a) + strconv.Itoa(b)
	bf := strconv.Itoa(b) + strconv.Itoa(a)
	afn, _ := strconv.Atoi(af)
	bfn, _ := strconv.Atoi(bf)
	if afn < bfn {
		return true
	}
	return false
}
