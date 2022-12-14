package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	Src1 string
	Src2 string
}

func Test_searchInsert1(t *testing.T) {
	var test []Test
	test = append(test, Test{Src1: "a", Src2: "aa"})
	test = append(test, Test{Src1: "b", Src2: "bb"})
	res := make(map[string]string, len(test))
	for _, v := range test {
		res[v.Src1] = v.Src2
	}
	fmt.Println(res)
}
func Test_searchInsert(t *testing.T) {
	assert.EqualValues(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))
	assert.EqualValues(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))
	assert.EqualValues(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))
	assert.EqualValues(t, 0, searchInsert([]int{1, 3, 5, 6}, 0))
}
