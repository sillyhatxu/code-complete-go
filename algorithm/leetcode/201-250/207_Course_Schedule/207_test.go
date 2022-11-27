package _07_Course_Schedule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canFinish(t *testing.T) {
	assert.EqualValues(t, true, canFinish(5, [][]int{{1, 0}, {2, 1}, {3, 2}, {4, 3}, {3, 1}}))
}
