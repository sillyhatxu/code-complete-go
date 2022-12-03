package _71_Sum_of_Two_Integers

func getSum(a int, b int) int {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	}
	for b != 0 {
		a, b = a^b, a&b<<1
	}
	return a
}
