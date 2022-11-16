package _50_Evaluate_Reverse_Polish_Notation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_evalRPN(t *testing.T) {
	assert.EqualValues(t, 22, evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
