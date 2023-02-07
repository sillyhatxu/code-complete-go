package v2

type NestedIterator struct {
	id   int
	list []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	ni := &NestedIterator{
		id:   0,
		list: make([]int, 0, 250),
	}

	ni.add(nestedList)

	return ni
}

func (this *NestedIterator) add(nestedList []*NestedInteger) {
	for _, ns := range nestedList {
		if ns.IsInteger() {
			this.list = append(this.list, ns.GetInteger())
		} else {
			this.add(ns.GetList())
		}
	}
}

func (this *NestedIterator) Next() int {
	res := this.list[this.id]
	this.id++
	return res
}

func (this *NestedIterator) HasNext() bool {
	return this.id < len(this.list)
}

/** --------------------------- 分割线 --------------------------- **/
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}
