package parking_dilemma

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_carParkingRoof(t *testing.T) {
	assert.EqualValues(t, 6, carParkingRoof([]int64{6, 2, 12, 7}, 3))
	assert.EqualValues(t, 9, carParkingRoof([]int64{2, 10, 8, 17}, 3))
}
