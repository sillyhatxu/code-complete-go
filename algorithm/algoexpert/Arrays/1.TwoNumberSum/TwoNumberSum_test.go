package TwoNumberSum

import (
	"testing"
)

func TestTwoNumberSum(t *testing.T) {
	result := TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)
	expect := []int{11, -1}
	if result[0] != expect[0] {
		t.Fatalf("%v != %v", result, expect)
	}
	if result[1] != expect[1] {
		t.Fatalf("%v != %v", result, expect)
	}
}
