// https://leetcode.com/problems/pascals-triangle/
package main

func generate(numRows int) [][]int {
	rv := [][]int{}
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = rv[i-1][j-1] + rv[i-1][j]
		}
		rv = append(rv, row)
	}
	return rv
}
