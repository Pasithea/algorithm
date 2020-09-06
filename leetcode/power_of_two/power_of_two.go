// https://leetcode.com/problems/power-of-two/submissions/
package main

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 2 {
		return true
	}
	if x := (n - 1) & n; x == 0 {
		return true
	}
	return false
}
