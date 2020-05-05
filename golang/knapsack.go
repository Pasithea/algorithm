package main

import (
	"fmt"
)

func main() {
	item := []int{3, 5, 7}
	fmt.Println(knapsack(item, 9))
}

func knapsack(item []int, w int) int {
	length := len(item)
	states := make([]bool, w+1)
	states[0] = true
	if item[0] <= w {
		states[item[0]] = true
	}
	for i := 1; i < length; i++ {
		for j := w - item[i]; j >= 0; j-- {
			if states[j] {
				states[j+item[i]] = true
			}
		}
	}
	for j := w; j >= 0; j-- {
		if states[j] {
			return j
		}
	}
	return 0
}
