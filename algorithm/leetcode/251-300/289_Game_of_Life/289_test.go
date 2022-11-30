package _89_Game_of_Life

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_gameOfLife_test(t *testing.T) {
	fmt.Println(0 & 1)
	fmt.Println(1 & 1)
	fmt.Println(2 & 1)
	fmt.Println(3 & 1)
	fmt.Println("-----")
	fmt.Println(1 >> 1)
	fmt.Println(3 >> 1)
	fmt.Println(1 >> 1)
	fmt.Println(2 >> 1)
	fmt.Println(0 >> 1)
	fmt.Println("-----")
	fmt.Println(1 & 1)
	fmt.Println(3 & 1)
	fmt.Println(1 & 1)
	fmt.Println(2 & 1)
	fmt.Println(0 & 1)
}
func Test_gameOfLife(t *testing.T) {
	input := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	expected := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}
	gameOfLife(input)
	assert.EqualValues(t, expected, input)
}
