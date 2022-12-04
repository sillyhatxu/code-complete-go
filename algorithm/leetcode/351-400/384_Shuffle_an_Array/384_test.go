package _84_Shuffle_an_Array

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
