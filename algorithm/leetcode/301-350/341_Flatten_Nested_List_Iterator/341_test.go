package _41_Flatten_Nested_List_Iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NestedIterator1(t *testing.T) {
	val1, val2, val3, val4, val5 := 1, 2, 3, 4, 5
	value1 := &NestedInteger{value: &val1, array: nil}
	value2 := &NestedInteger{value: &val2, array: nil}
	value3 := &NestedInteger{value: &val3, array: nil}
	value4 := &NestedInteger{value: &val4, array: nil}
	value5 := &NestedInteger{value: &val5, array: nil}
	var input []*NestedInteger
	input = append(input, &NestedInteger{value: nil, array: []*NestedInteger{value1, value2}})
	input = append(input, &NestedInteger{value: value3.value, array: nil})
	input = append(input, &NestedInteger{value: nil, array: []*NestedInteger{value4, value5}})
	var res []int
	nestedIterator := Constructor(input)
	for nestedIterator.HasNext() {
		res = append(res, nestedIterator.Next())
	}
	assert.EqualValues(t, []int{1, 2, 3, 4, 5}, res)
}
