// https://leetcode.com/problems/climbing-stairs/
package main

var m map[int]int = make(map[int]int)

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	if _, ok := m[n]; ok {
	} else {
		m[n] = climbStairs(n-1) + climbStairs(n-2)
	}
	return m[n]
}
