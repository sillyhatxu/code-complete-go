package main

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = x * -1
	}
	res := 0
	src := strconv.Itoa(x)
	for i := len(src) - 1; i >= 0; i-- {
		res = res*10 + int(src[i]-'0')
	}
	if res > math.MaxInt32 || res < -math.MaxInt32-1 {
		return 0
	}
	return res * sign
}

func reverse1(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = x * -1
	}
	var result int64 = 0
	src := strconv.Itoa(x)
	for i := len(src) - 1; i >= 0; i-- {
		r := int64(x % 10)
		for j := 0; j < i; j++ {
			r = r * 10
		}
		result += r
		x = x / 10
	}
	if result > math.MaxInt32 || result < -math.MaxInt32-1 {
		return 0
	}
	return int(result) * sign
}
