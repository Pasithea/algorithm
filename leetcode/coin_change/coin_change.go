// https://leetcode.com/problems/coin-change/
package main

func coinChange(coins []int, amount int) int {
	count := make([]int, amount+1)
	count[0] = 0
	for i := 1; i < amount+1; i++ {
		count[i] = amount + 1
		for _, coin := range coins {
			if i-coin >= 0 {
				count[i] = min(count[i-coin]+1, count[i])
			}
		}
	}
	ans := count[amount]
	if ans == amount+1 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
