package ValidateSubsequence

import "testing"

func TestIsValidSubsequence(t *testing.T) {
	result := false
	//result = IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10})
	//if !result {
	//	t.Fatalf("%v", result)
	//}
	//result = IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 25, 6, -1, 8, 10})
	//if !result {
	//	t.Fatalf("%v", result)
	//}
	//result = IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 22, 25, 6, -1, 8, 10})
	//if result {
	//	t.Fatalf("%v", result)
	//}
	result = IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 25, 22, 6, -1, 8, 10})
	if result {
		t.Fatalf("%v", result)
	}
}
