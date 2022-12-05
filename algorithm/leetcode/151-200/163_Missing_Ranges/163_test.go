package _63_Missing_Ranges

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMissingRanges(t *testing.T) {
	result := findMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99)
	fmt.Println(result)
	assert.EqualValues(t, "2", result[0])
	assert.EqualValues(t, "4->49", result[1])
	assert.EqualValues(t, "51->74", result[2])
	assert.EqualValues(t, "76->99", result[3])
}
