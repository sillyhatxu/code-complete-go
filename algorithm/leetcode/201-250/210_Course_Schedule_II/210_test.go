package _10_Course_Schedule_II

import (
	"fmt"
	"testing"
)

func Test_findOrder(t *testing.T) {
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 1}, {2, 0}, {0, 3}}))
}
