package main


type MinStack struct {
	data   []int
	min    int
	length int
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x - this.min)
	this.length++
	if this.length == 1 {
		this.min = x
	} else if x <= this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	x := this.data[this.length-1]
	this.data = this.data[:len(this.data)-1]
	this.length--
	if x <= 0 {
		this.min = this.min - x
	}
}

func (this *MinStack) GetMin() int {
	return this.min
}

func (this *MinStack) Top() int {
	top := this.data[this.length-1]
	if this.length == 1 {
		return top
	} else if top <= 0 {
		return this.min - top
	} else {
		return top + this.min
	}
}
