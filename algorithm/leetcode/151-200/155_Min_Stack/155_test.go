package _55_Min_Stack

import (
	"fmt"
	"testing"
)

func Test_MinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}

func Test_MinStack1(t *testing.T) {
	minStack := Constructor()
	minStack.Push(0)
	minStack.Push(1)
	minStack.Push(0)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
}

func Test_MinStack2(t *testing.T) {
	minStack := Constructor()
	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)
	fmt.Println(minStack.Top())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	minStack.Push(2147483647)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	minStack.Push(-2147483648)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
}
