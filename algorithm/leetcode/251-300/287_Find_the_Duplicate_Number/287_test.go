package _87_Find_the_Duplicate_Number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	assert.EqualValues(t, 2, findDuplicate([]int{1, 3, 4, 2, 2}))
}
