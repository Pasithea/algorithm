// https://leetcode.com/problems/power-of-three/
package main

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%3 != 0 {
		return false
	}
	for n != 0 {
		n /= 3
		if n%3 != 0 && n != 1 {
			return false
		}
	}
	return true
}
