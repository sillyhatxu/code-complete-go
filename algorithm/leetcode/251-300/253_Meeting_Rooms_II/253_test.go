package _53_Meeting_Rooms_II

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minMeetingRooms(t *testing.T) {
	assert.EqualValues(t, 1, minMeetingRooms([][]int{{13, 15}, {1, 13}}))
	assert.EqualValues(t, 2, minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
	assert.EqualValues(t, 1, minMeetingRooms([][]int{{7, 10}, {2, 4}}))
}
