package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	assert.EqualValues(t, 1, climbStairs(1))
	assert.EqualValues(t, 2, climbStairs(2))
	assert.EqualValues(t, 3, climbStairs(3))
	assert.EqualValues(t, 5, climbStairs(4))
	assert.EqualValues(t, 8, climbStairs(5))
	assert.EqualValues(t, 13, climbStairs(6))
	assert.EqualValues(t, 21, climbStairs(7))
	assert.EqualValues(t, 34, climbStairs(8))
	assert.EqualValues(t, 55, climbStairs(9))
	assert.EqualValues(t, 89, climbStairs(10))
	assert.EqualValues(t, 144, climbStairs(11))
	assert.EqualValues(t, 267914296, climbStairs(41))
	assert.EqualValues(t, 14930352, climbStairs(35))
}

func Test_climbStairs3(t *testing.T) {
	assert.EqualValues(t, 1, climbStairs3(1))
	assert.EqualValues(t, 2, climbStairs3(2))
	assert.EqualValues(t, 3, climbStairs3(3))
	assert.EqualValues(t, 5, climbStairs3(4))
	assert.EqualValues(t, 8, climbStairs3(5))
	assert.EqualValues(t, 13, climbStairs3(6))
	assert.EqualValues(t, 21, climbStairs3(7))
	assert.EqualValues(t, 34, climbStairs3(8))
	assert.EqualValues(t, 55, climbStairs3(9))
	assert.EqualValues(t, 89, climbStairs3(10))
	assert.EqualValues(t, 144, climbStairs3(11))
	assert.EqualValues(t, 267914296, climbStairs3(41))
	assert.EqualValues(t, 14930352, climbStairs3(35))
}

func Test_climbStairs1(t *testing.T) {
	assert.EqualValues(t, 1, climbStairs1(1))
	assert.EqualValues(t, 2, climbStairs1(2))
	assert.EqualValues(t, 3, climbStairs1(3))
	assert.EqualValues(t, 5, climbStairs1(4))
	assert.EqualValues(t, 8, climbStairs1(5))
	assert.EqualValues(t, 13, climbStairs1(6))
	assert.EqualValues(t, 21, climbStairs1(7))
	assert.EqualValues(t, 34, climbStairs1(8))
	assert.EqualValues(t, 55, climbStairs1(9))
	assert.EqualValues(t, 89, climbStairs1(10))
	assert.EqualValues(t, 144, climbStairs1(11))
	assert.EqualValues(t, 267914296, climbStairs1(41))
	assert.EqualValues(t, 14930352, climbStairs1(35))
}

func Test_climbStairsOptimize(t *testing.T) {
	assert.EqualValues(t, 1, climbStairsOptimize(1))
	assert.EqualValues(t, 2, climbStairsOptimize(2))
	assert.EqualValues(t, 3, climbStairsOptimize(3))
	assert.EqualValues(t, 5, climbStairsOptimize(4))
	assert.EqualValues(t, 8, climbStairsOptimize(5))
	assert.EqualValues(t, 13, climbStairsOptimize(6))
	assert.EqualValues(t, 21, climbStairsOptimize(7))
	assert.EqualValues(t, 34, climbStairsOptimize(8))
	assert.EqualValues(t, 55, climbStairsOptimize(9))
	assert.EqualValues(t, 89, climbStairsOptimize(10))
	assert.EqualValues(t, 144, climbStairsOptimize(11))
	assert.EqualValues(t, 267914296, climbStairsOptimize(41))
	assert.EqualValues(t, 14930352, climbStairsOptimize(35))
}
