// https://leetcode.com/problems/power-of-four/
package main

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}
	if x, y := (num-1)&num, num&0x55555555; x == 0 && y == num {
		return true
	}
	return false
}
