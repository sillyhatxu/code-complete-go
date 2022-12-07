package _51_Flatten_2D_Vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Vector2D(t *testing.T) {
	iterator := Constructor([][]int{{1, 2}, {3}, {4}})
	assert.EqualValues(t, 1, iterator.Next())
	assert.EqualValues(t, 2, iterator.Next())
	assert.EqualValues(t, 3, iterator.Next())
	assert.EqualValues(t, true, iterator.HasNext())
	assert.EqualValues(t, true, iterator.HasNext())
	assert.EqualValues(t, 4, iterator.Next())
	assert.EqualValues(t, false, iterator.HasNext())
}
