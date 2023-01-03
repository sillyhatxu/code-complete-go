package _18_Pascal_s_Triangle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generate(t *testing.T) {
	assert.EqualValues(t, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}, generate(5))
	//fmt.Println(0 >> 1)
	//fmt.Println(1 >> 1)
	//fmt.Println(2 >> 1)
	//fmt.Println(3 >> 1)
	//fmt.Println(4 >> 1)
}
