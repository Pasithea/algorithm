// https://leetcode.com/problems/kth-largest-element-in-an-array/
package main

func findKthLargest(nums []int, k int) int {
	heap := make([]int, 0, k)
	for i := 0; i < len(nums); i++ {
		if len(heap) < k {
			heap = append(heap, nums[i])
			up(heap, len(heap)-1)
		} else {
			if nums[i] > heap[0] {
				heap[0] = nums[i]
				down(heap, 0, len(heap)-1)
			}
		}
	}
	return heap[0]
}

func up(heap []int, n int) {
	for n > 0 {
		p := (n - 1) / 2
		if n == p {
			return
		}
		if heap[n] < heap[p] {
			heap[p], heap[n] = heap[n], heap[p]
		} else {
			return
		}
		n = p
	}
}

func down(heap []int, n, k int) {
	for n <= k {
		c := n
		c1 := c*2 + 1
		c2 := c1 + 1
		if c1 <= k && heap[c1] < heap[c] {
			c = c1
		}
		if c2 <= k && heap[c2] < heap[c] {
			c = c2
		}
		if c == n {
			return
		}
		heap[c], heap[n] = heap[n], heap[c]
		n = c
	}
}
