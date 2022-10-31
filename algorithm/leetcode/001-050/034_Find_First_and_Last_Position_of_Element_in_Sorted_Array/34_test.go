package _34_Find_First_and_Last_Position_of_Element_in_Sorted_Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchRange(t *testing.T) {
	//result1 := searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	//assert.EqualValues(t, []int{-1, -1}, result1)
	result2 := searchRange([]int{2, 2}, 2)
	assert.EqualValues(t, []int{0, 1}, result2)
	//result3 := search([]int{3, 1}, 1)
	//assert.EqualValues(t, 1, result3)
	//result4 := search([]int{3, 1}, 3)
	//assert.EqualValues(t, 0, result4)
	//result5 := search([]int{3, 4, 5, 6, 1, 2}, 2)
	//assert.EqualValues(t, 5, result5)
	//result6 := search([]int{1, 3}, 0)
	//assert.EqualValues(t, -1, result6)
}
