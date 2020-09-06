// https://leetcode.com/problems/top-k-frequent-elements/
package main

type counter struct {
	num   int
	count int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]*counter)
	for i := 0; i < len(nums); i++ {
		if x, ok := m[nums[i]]; ok {
			x.count++
		} else {
			m[nums[i]] = &counter{num: nums[i], count: 1}
		}
	}
	heap := make([]*counter, 0, k)
	for _, v := range m {
		if len(heap) < k {
			heap = append(heap, v)
			up(heap, len(heap)-1)
		} else {
			if v.count > heap[0].count {
				heap[0] = v
				down(heap, 0, len(heap)-1)
			}
		}
	}
	res := make([]int, k, k)
	for i := 0; i < k; i++ {
		res[i] = heap[i].num
	}
	return res
}

func up(heap []*counter, n int) {
	for n > 0 {
		p := (n - 1) / 2
		if p == n {
			return
		}
		if heap[n].count < heap[p].count {
			heap[p], heap[n] = heap[n], heap[p]
		} else {
			return
		}
		n = p
	}
}

func down(heap []*counter, n, k int) {
	for n <= k {
		c := n
		c1 := c*2 + 1
		c2 := c1 + 1
		if c1 <= k && heap[c1].count < heap[c].count {
			c = c1
		}
		if c2 <= k && heap[c2].count < heap[c].count {
			c = c2
		}
		if n == c {
			return
		}
		heap[c], heap[n] = heap[n], heap[c]
		n = c
	}
}
