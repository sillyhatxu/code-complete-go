package parking_dilemma

import (
	"sort"
)

func carParkingRoof(cars []int64, k int32) int64 {
	sort.Slice(cars, func(i, j int) bool {
		return cars[i] < cars[j]
	})
	minDist := cars[k-1] - cars[0] + 1
	for i := k; i <= int32(len(cars)); i++ {
		dist := cars[i-1] - cars[i-k]
		if dist < minDist {
			minDist = dist + 1
		}
	}
	return minDist
}
