package _07_Course_Schedule

import "testing"

func Test_canFinish(t *testing.T) {
	canFinish(5, [][]int{{1, 0}, {2, 1}, {3, 2}, {4, 3}, {3, 1}})
}
