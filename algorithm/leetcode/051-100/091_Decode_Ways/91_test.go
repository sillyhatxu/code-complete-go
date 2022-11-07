package _91_Decode_Ways

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numDecodings1(t *testing.T) {
	i, j := 1, 3
	i, j = j, i+j
	//i = 3; j = 4
	fmt.Println(i, "---", j)
}
func Test_numDecodings(t *testing.T) {
	//result1 := numDecodings("226")
	//assert.EqualValues(t, 3, result1)
	result2 := numDecodings("12101226")
	assert.EqualValues(t, 10, result2)
	//result3 := numDecodings("1210")
	//assert.EqualValues(t, 2, result3)
}
