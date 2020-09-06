// https://leetcode.com/problems/min-stack/
package main

type MinStack struct {
	stack  []int
	curMin *Min
}

type Min struct {
	pos    int
	preMin *Min
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: []int{}, curMin: nil}
}

func (ms *MinStack) Push(x int) {
	ms.stack = append(ms.stack, x)
	if ms.curMin == nil {
		ms.curMin = &Min{pos: len(ms.stack) - 1, preMin: nil}
	} else {
		if ms.stack[ms.curMin.pos] > x {
			preMin := ms.curMin
			ms.curMin = &Min{pos: len(ms.stack) - 1, preMin: preMin}
		}
	}
}

func (ms *MinStack) Pop() {
	pos := len(ms.stack) - 1
	if pos < 0 {
		return
	}
	if pos == ms.curMin.pos {
		ms.curMin = ms.curMin.preMin
	}
	ms.stack = ms.stack[:pos]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.stack[ms.curMin.pos]
}
