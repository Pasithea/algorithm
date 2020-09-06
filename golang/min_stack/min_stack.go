package main

type MinStack struct {
	data   []int
	min    int
	length int
}

func (ms *MinStack) Push(x int) {
	ms.data = append(ms.data, x-ms.min)
	ms.length++
	if ms.length == 1 {
		ms.min = x
	} else if x <= ms.min {
		ms.min = x
	}
}

func (ms *MinStack) Pop() {
	x := ms.data[ms.length-1]
	ms.data = ms.data[:len(ms.data)-1]
	ms.length--
	if x <= 0 {
		ms.min = ms.min - x
	}
}

func (ms *MinStack) GetMin() int {
	return ms.min
}

func (ms *MinStack) Top() int {
	top := ms.data[ms.length-1]
	if ms.length == 1 {
		return top
	} else if top <= 0 {
		return ms.min - top
	} else {
		return top + ms.min
	}
}
