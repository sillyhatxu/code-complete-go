package _41_Flatten_Nested_List_Iterator

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */
//TODO
type NestedIterator struct {
	Stack Stack
}

/*
Custom type Stack. Start...
*/
type Stack []Element

type Element struct {
	data []*NestedInteger
	i    int
}

func (s *Stack) Top() Element {
	if len(*s) == 0 {
		return Element{}
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() Element {
	top := s.Top()
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *Stack) Push(x Element) {
	*s = append(*s, x)
}

func (s *Stack) Len() int {
	return len(*s)
}

/*
Custom type Stack. End...
*/

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	res := &NestedIterator{Stack: make(Stack, 0)}
	if len(nestedList) == 0 {
		return res
	}
	res.Stack.Push(Element{data: nestedList, i: 0})
	return res
}

func (this *NestedIterator) Next() int {
	top := this.Stack.Pop()
	if top.data[top.i].IsInteger() {
		if top.i < len(top.data)-1 {
			this.Stack.Push(Element{top.data, top.i + 1})
		}
		return top.data[top.i].GetInteger()
	} else {
		if top.i < len(top.data)-1 {
			this.Stack.Push(Element{top.data, top.i + 1})
		}
		nextList := top.data[top.i].GetList()
		this.Stack.Push(Element{nextList, 0})
		return this.Next()
	}
}

func (this *NestedIterator) HasNext() bool {
	top := this.Stack.Top()
	for this.Stack.Len() > 0 && !top.data[top.i].IsInteger() {
		top = this.Stack.Pop()
		if top.i < len(top.data)-1 {
			this.Stack.Push(Element{top.data, top.i + 1})
		}
		nextList := top.data[top.i].GetList()
		if len(nextList) != 0 {
			this.Stack.Push(Element{nextList, 0})
		}
		top = this.Stack.Top()
	}
	return this.Stack.Len() > 0
}

type NestedInteger struct {
	value *int
	array []*NestedInteger
}

func (this NestedInteger) IsInteger() bool {
	return this.value != nil
}

func (this NestedInteger) GetInteger() int {
	return *this.value
}

func (this *NestedInteger) SetInteger(value int) {
	this.value = &value
}

func (this *NestedInteger) Add(elem NestedInteger) {

}

func (this NestedInteger) GetList() []*NestedInteger {
	return this.array
}
