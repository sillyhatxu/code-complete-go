package _51_Flatten_2D_Vector

type Vector2D struct {
	Node *Node
}

type Node struct {
	Val  int
	Next *Node
}

func Constructor(vec [][]int) Vector2D {
	temp := &Node{}
	current := temp
	for i := 0; i < len(vec); i++ {
		for j := 0; j < len(vec[i]); j++ {
			current.Next = &Node{Val: vec[i][j]}
			current = current.Next
		}
	}
	return Vector2D{Node: temp}
}

func (this *Vector2D) Next() int {
	this.Node = this.Node.Next
	return this.Node.Val
}

func (this *Vector2D) HasNext() bool {
	return this.Node.Next != nil
}

/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(vec);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
