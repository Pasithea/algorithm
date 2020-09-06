// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
package main

func maxProfit(prices []int) int {
	total := 0
	length := len(prices)
	if length < 2 {
		return total
	}
	for i := 0; i < length-1; i++ {
		diff := prices[i+1] - prices[i]
		if diff > 0 {
			total += diff
		}
	}
	return total
}
