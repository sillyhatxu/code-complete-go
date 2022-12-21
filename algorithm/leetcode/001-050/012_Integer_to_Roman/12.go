package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return fmt.Sprintf("%s%s%s%s", M[num/1000], C[num/100%10], X[num/10%10], I[num%10])
}

func intToRomanOptimize(num int) string {
	m, n, s, result := []string{"I", "V", "X", "L", "C", "D", "M", "", ""}, 1000, 3, ""
	for n > 0 {
		result += numToRoman(num/n, m[s*2], m[s*2+1], m[s*2+2])
		num -= num / n * n
		n = n / 10
		s--
	}
	return result
}

func numToRoman(num int, one string, five string, ten string) string {
	switch {
	case num == 9:
		return one + ten
	case num >= 5:
		return five + strings.Repeat(one, num-5)
	case num == 4:
		return one + five
	case num < 4:
		return strings.Repeat(one, num)
	default:
		return ""
	}
}
