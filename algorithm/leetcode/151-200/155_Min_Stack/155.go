package _55_Min_Stack

import "math"

type MinStack struct {
	node *Node
	min  int
}
type Node struct {
	Val  int
	Next *Node
}

func Constructor() MinStack {
	return MinStack{min: math.MaxInt64}
}

func (this *MinStack) Push(val int) {
	current := &Node{Val: val, Next: &Node{Val: val}}
	if this.node == nil {
		this.node = current
		this.min = val
		return
	}
	if val <= this.min {
		preMin := &Node{Val: this.min, Next: this.node}
		current.Next = preMin
		this.min = val
	} else {
		current.Next = this.node
	}
	this.node = current
}

func (this *MinStack) Pop() {
	if this.node == nil {
		return
	}
	if this.node.Val == this.min {
		this.node = this.node.Next
		this.min = this.node.Val
	}
	this.node = this.node.Next

}

func (this *MinStack) Top() int {
	if this.node == nil {
		return math.MinInt64
	}
	return this.node.Val
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
