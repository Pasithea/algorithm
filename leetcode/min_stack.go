// https://leetcode.com/problems/min-stack/


type MinStack struct {
	stack []int
	curMin *Min
}

type Min struct {
	pos int
	preMin *Min
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: []int{}, curMin: nil}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if this.curMin == nil {
		this.curMin = &Min{pos: len(this.stack)-1, preMin: nil}
	} else {
		if this.stack[this.curMin.pos] > x {
			preMin := this.curMin
			this.curMin = &Min{pos: len(this.stack)-1, preMin: preMin}
		}
	}
}


func (this *MinStack) Pop()  {
	pos := len(this.stack) - 1
	if pos < 0 {
		return
	}
	if pos == this.curMin.pos {
		this.curMin = this.curMin.preMin
	}
	this.stack = this.stack[:pos]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.stack[this.curMin.pos]
}
