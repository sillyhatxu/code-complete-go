package parking_dilemma

import (
	"sort"
)

func carParkingRoof(cars []int64, k int32) int64 {
	sort.Slice(cars, func(i, j int) bool {
		return cars[i] < cars[j]
	})
	size := len(cars)
	carParking := func(dist int64) bool {
		count := int64(1)
		prev := cars[0]
		for i := 1; i < size; i++ {
			if cars[i]-prev >= dist {
				count++
				prev = cars[i]
				if count > int64(k) {
					return false
				}
			}
		}
		return true
	}
	left, right := int64(1), cars[size-1]-cars[0]
	for left < right {
		mid := left + (right-left)/2
		if carParking(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
