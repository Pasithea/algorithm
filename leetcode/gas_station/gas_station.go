// https://leetcode.com/problems/gas-station/
package main

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	station := 0
	curTank, totalTank := 0, 0
	for i := 0; i < n; i++ {
		totalTank += gas[i] - cost[i]
		curTank += gas[i] - cost[i]
		if curTank < 0 {
			station = i + 1
			curTank = 0
		}
	}
	if totalTank >= 0 {
		return station
	}
	return -1
}
