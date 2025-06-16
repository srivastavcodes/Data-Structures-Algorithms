package main

type MyQueue struct {
	input  []int
	output []int
	peekEl int
}

func Constructor() MyQueue {
	return MyQueue{
		input:  make([]int, 0),
		output: make([]int, 0),
		peekEl: -1,
	}
}

func (mq *MyQueue) Push(x int) {
	if len(mq.input) == 0 {
		mq.peekEl = x
	}
	mq.input = append(mq.input, x)
}

func (mq *MyQueue) Pop() int {
	if len(mq.output) == 0 {
		// push input to output
		for len(mq.input) > 0 {
			topIdx := len(mq.input) - 1
			mq.output = append(mq.output, mq.input[topIdx])
			mq.input = mq.input[:topIdx]
		}
	}
	value := mq.output[len(mq.output)-1]
	mq.output = mq.output[:len(mq.output)-1]
	return value
}

func (mq *MyQueue) Peek() int {
	if len(mq.output) == 0 {
		return mq.peekEl
	}
	return mq.output[len(mq.output)-1]
}

func (mq *MyQueue) Empty() bool {
	return len(mq.input) == 0 && len(mq.output) == 0
}
