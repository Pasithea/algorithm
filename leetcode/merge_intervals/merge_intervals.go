// https://leetcode.com/problems/merge-intervals/
package main

func merge(intervals [][]int) [][]int {
	sort(intervals)
	merged := make([][]int, 0, len(intervals))
	for _, interval := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			if merged[len(merged)-1][1] < interval[1] {
				merged[len(merged)-1][1] = interval[1]
			}
		}
	}
	return merged
}

func sort(x [][]int) {
	if len(x) < 2 {
		return
	}
	qsort(x, 0, len(x)-1)
}

func qsort(x [][]int, start, end int) {
	if start >= end {
		return
	}
	i := partition(x, start, end)
	qsort(x, start, i-1)
	qsort(x, i+1, end)
}

func partition(x [][]int, start, end int) int {
	pviot := x[end][0]
	i := start
	for j := start; j < end; j++ {
		if x[j][0] < pviot {
			x[j], x[i] = x[i], x[j]
			i++
		}
	}
	x[i], x[end] = x[end], x[i]
	return i
}
