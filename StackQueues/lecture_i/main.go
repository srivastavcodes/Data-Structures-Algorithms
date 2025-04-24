package main

import "fmt"

func main() {
	stack := NewStack()
	err := stack.Push(10)
	err = stack.Push(20)
	err = stack.Push(30)
	err = stack.Push(40)
	err = stack.Push(50)
	err = stack.Push(60)
	err = stack.Push(70)
	err = stack.Push(80)
	err = stack.Push(90)
	err = stack.Push(100)
	err = stack.Push(110)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(stack.Top())
	fmt.Println(stack.Size())
}
