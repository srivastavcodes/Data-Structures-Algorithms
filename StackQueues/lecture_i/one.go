package main

type Stack struct {
	top  int
	data []int
}

func NewStack() *Stack {
	return &Stack{top: -1, data: make([]int, 10)}
}

func (stk *Stack) Push(val int) {
	stk.top++
	if stk.top >= cap(stk.data) {
		stk.top--
	}
	stk.data[stk.top] = val
}

func (stk *Stack) Top() int {
	if stk.top < 0 {
		return -1
	}
	return stk.data[stk.top]
}

func (stk *Stack) Pop() int {
	stk.top--
	if stk.top < 0 {
		stk.top++
		return -1
	}
	return stk.data[stk.top]
}

func (stk *Stack) Size() int {
	if stk.top < 0 {
		return -1
	}
	return len(stk.data)
}
