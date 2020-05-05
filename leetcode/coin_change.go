// https://leetcode.com/problems/coin-change/


func coinChange(coins []int, amount int) int {
	count := make([]int, amount+1)
	count[0] = 0
	for i := 1; i < amount + 1; i++ {
		count[i] = amount + 1 // as max
	}
	for i := 0; i < amount + 1; i++ {
		for _, coin := range coins {
			if i - coin < 0 {
				continue
			}
			count[i] = min(count[i], count[i-coin]+1)
		}
	}
	ans := count[amount]
	if ans == amount + 1 {
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
