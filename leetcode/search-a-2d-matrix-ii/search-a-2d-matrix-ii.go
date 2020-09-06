// https://leetcode.com/problems/search-a-2d-matrix-ii/
package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		x := matrix[i][j]
		if x == target {
			return true
		} else if x < target {
			i++
		} else {
			j--
		}
	}
	return false
}
