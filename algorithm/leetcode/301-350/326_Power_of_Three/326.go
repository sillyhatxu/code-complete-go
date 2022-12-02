package _26_Power_of_Three

import (
	"math"
)

//non-loops
func isPowerOfThree(n int) bool {
	//temp := math.Log10(float64(n)) / math.Log10(float64(3))
	//_, res := math.Modf(temp)
	//return res == 0.0
	if n <= 0 {
		return false
	}
	quo, _ := math.Modf(math.Log10(float64(n)) / math.Log10(3))
	return n == int(math.Pow(3, quo))
}

func isPowerOfThreeError1(n int) bool {
	//编译问题 (operator % is not defined on untyped int)
	//return (math.Log10(float64(n)) / math.Log10(float64(3))) % 1 == 0
	return false
}

func isPowerOfThreeError2(n int) bool {
	//精度问题; n=27 temp = 3.0000000000000004
	temp := math.Log10(float64(n)) / math.Log10(float64(3))
	_, res := math.Modf(temp)
	return res == 0.0
}

//loops
func isPowerOfThree1(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func isPowerOfThree2(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 || n == 3 {
		return true
	}
	current := 3
	for i := 1; i < n; i++ {
		current *= 3
		if current == n {
			return true
		}
		if current > n {
			break
		}
	}
	return false
}
