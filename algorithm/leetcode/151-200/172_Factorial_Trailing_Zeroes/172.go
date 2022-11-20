package _72_Factorial_Trailing_Zeroes

func trailingZeroes(n int) int {
	numOf5 := 0
	for n > 4 {
		numOf5, n = numOf5+n/5, n/5
	}
	return numOf5
}
