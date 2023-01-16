package _384_Shuffle_an_Array_2

import (
	"fmt"
	"testing"
)

func Test_Solution(t *testing.T) {
	solution := Constructor([]int{1, 2, 3})
	fmt.Println(solution.Shuffle())
	fmt.Println(solution.Reset())
	fmt.Println(solution.Shuffle())
}
