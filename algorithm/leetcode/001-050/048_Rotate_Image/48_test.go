package _48_Rotate_Image

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(input)
	fmt.Println(input)
}
