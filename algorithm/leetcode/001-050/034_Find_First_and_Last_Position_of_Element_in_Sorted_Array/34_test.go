package _34_Find_First_and_Last_Position_of_Element_in_Sorted_Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchRange(t *testing.T) {
	//result1 := searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	//assert.EqualValues(t, []int{-1, -1}, result1)
	result2 := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	assert.EqualValues(t, []int{3, 4}, result2)
}
