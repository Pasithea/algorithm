// https://leetcode.com/problems/search-a-2d-matrix/
package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	start, end := 0, m*n-1
	for start <= end {
		mid := (start + end) / 2
		posM := mid / n
		posN := mid % n
		if matrix[posM][posN] > target {
			end = mid - 1
		} else if matrix[posM][posN] < target {
			start = mid + 1
		} else {
			return true
		}
	}
	return false
}
