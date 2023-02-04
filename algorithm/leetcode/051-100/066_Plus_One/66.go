package main

func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] == 10 {
			digits[i] = 0
			digits[i-1]++
		}
	}
	if digits[0] == 10 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}
	return digits
}

func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= -1; i-- {
		if i == -1 {
			digits[0] = 0
			digits = append([]int{1}, digits...)
			break
		}
		digits[i]++
		if digits[i] != 10 {
			break
		}
		digits[i] = 0 // digits[i] = 10
	}
	return digits
}

func plusOne1(digits []int) []int {
	temp := digits[len(digits)-1] + 1
	if temp < 10 {
		digits[len(digits)-1] = temp
		return digits
	}
	carry := 1
	digits[len(digits)-1] = 0
	for i := len(digits) - 2; i >= 0; i-- {
		temp = digits[i] + carry
		if temp < 10 {
			digits[i] = temp
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append(digits[:0], append([]int{1}, digits[0:]...)...)
}
