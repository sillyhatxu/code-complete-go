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

type NestedIterator struct {
	nestedList   []*NestedInteger
	currentIndex int
	listIndex    int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{nestedList: nestedList, currentIndex: 0, listIndex: 0}
}

func (this *NestedIterator) Next() int {
	currentNestedInteger := this.nestedList[this.currentIndex]
	if currentNestedInteger.IsInteger() {
		this.currentIndex++
		return currentNestedInteger.GetInteger()
	}
	res := currentNestedInteger.GetList()[this.listIndex]
	this.listIndex++
	if this.listIndex == len(currentNestedInteger.GetList()) {
		this.listIndex = 0
		this.currentIndex++
	}
	return res.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	return this.nestedList != nil && this.nestedList[this.currentIndex] != nil
}

type NestedInteger struct {
	value *int
	next  []*NestedInteger
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
	this.next = append(this.next, &elem)
}

func (this NestedInteger) GetList() []*NestedInteger {
	return this.next
}
