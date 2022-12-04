package _80_Insert_Delete_GetRandom_O1

import (
	"fmt"
	"testing"
)

func Test_RandomizedSet(t *testing.T) {
	randomizedSet := Constructor()
	fmt.Println(randomizedSet.Insert(1))
	fmt.Println(randomizedSet.Remove(2))
	fmt.Println(randomizedSet.Insert(2))
	fmt.Println(randomizedSet.GetRandom())
	fmt.Println(randomizedSet.Remove(1))
	fmt.Println(randomizedSet.Insert(2))
	fmt.Println(randomizedSet.GetRandom())
}

func Test_RandomizedSet1(t *testing.T) {
	randomizedSet := Constructor()
	fmt.Println(randomizedSet.Insert(0))
	fmt.Println(randomizedSet.Insert(1))
	fmt.Println(randomizedSet.Remove(0))
	fmt.Println(randomizedSet.Insert(2))
	fmt.Println(randomizedSet.Remove(1))
	fmt.Println(randomizedSet.GetRandom())
}
