package main

import "fmt"

type Stack struct {
	top  int
	data []int
}

func NewStack() *Stack {
	return &Stack{top: -1, data: make([]int, 10)}
}

func (stk *Stack) Push(val int) error {
	stk.top++
	if stk.top >= cap(stk.data) {
		stk.top--
		return fmt.Errorf("stack overflow")
	}
	stk.data[stk.top] = val
	return nil
}

func (stk *Stack) Top() (int, error) {
	if stk.top < 0 {
		return -1, fmt.Errorf("stack underflow")
	}
	return stk.data[stk.top], nil
}

func (stk *Stack) Pop() (int, error) {
	stk.top--
	if stk.top < 0 {
		stk.top++
		return -1, fmt.Errorf("stack underflow")
	}
	return stk.data[stk.top], nil
}

func (stk *Stack) Size() (int, error) {
	if stk.top < 0 {
		return -1, fmt.Errorf("stack underflow")
	}
	return len(stk.data), nil
}
