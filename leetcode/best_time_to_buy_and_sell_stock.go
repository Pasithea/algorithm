// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/


func maxProfit(prices []int) int {
	min := 1 << 63 - 1
	maxValue := 0
	for i := 0; i < len(prices) - 1; i++ {
		if prices[i] < min {
			min = prices[i]
		}
		diff := prices[i+1] - min
		if diff > maxValue {
			maxValue = diff
		}
	}
	return maxValue
}
