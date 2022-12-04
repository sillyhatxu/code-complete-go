package _80_Insert_Delete_GetRandom_O1

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	valMap  map[int]int
	valList []int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{valMap: make(map[int]int), valList: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valMap[val]; ok {
		return false
	}
	index := len(this.valList)
	this.valMap[val] = index
	this.valList = append(this.valList, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.valMap[val]
	if !ok {
		return false
	}
	if index != len(this.valList)-1 {
		lastValue := this.valList[len(this.valList)-1]
		this.valMap[lastValue] = index
		this.valList[index] = lastValue
	}
	this.valList = this.valList[:len(this.valList)-1]
	delete(this.valMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.valList))
	return this.valList[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
