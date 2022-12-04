package _84_Shuffle_an_Array

import (
	"math/rand"
	"time"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	clone := make([]int, len(this.nums), len(this.nums))
	copy(clone, this.nums)
	rand.Seed(time.Now().UnixNano())
	for end := len(clone) - 1; end >= 0; end-- {
		randNum := rand.Intn(end + 1)
		clone[end], clone[randNum] = clone[randNum], clone[end]
	}
	return clone
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
