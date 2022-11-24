package _27_Basic_Calculator_II

import (
	"unicode"
)

func calculate(s string) int {
	var queue []int
	preSign := '+'
	num := 0
	for i, ch := range s {
		if unicode.IsDigit(ch) {
			num = (num * 10) + int(ch-'0')
			if i != len(s)-1 {
				continue
			}
		}
		if ch == ' ' && i != len(s)-1 {
			//Space or The last character must be a number
			continue
		}
		switch preSign {
		case '+':
			queue = append(queue, num)
		case '-':
			queue = append(queue, -num)
		case '*':
			queue[len(queue)-1] = queue[len(queue)-1] * num
		case '/':
			queue[len(queue)-1] = queue[len(queue)-1] / num
		}
		preSign = ch
		num = 0
	}
	res := 0
	for _, num := range queue {
		res += num
	}
	return res
}

func calculate1(s string) int {
	var stack []int
	preSign := '+'
	num := 0
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			num = (num * 10) + int(s[i]-'0')
			if i != len(s)-1 {
				continue
			}
		}
		if s[i] == ' ' && i != len(s)-1 {
			//Space or The last character must be a number
			continue
		}
		switch preSign {
		case '+':
			stack = append([]int{num}, stack...)
		case '-':
			stack = append([]int{-num}, stack...)
		case '*':
			stack[0] = stack[0] * num
		case '/':
			stack[0] = stack[0] / num
		}
		preSign = rune(s[i])
		num = 0
	}
	res := 0
	for _, num := range stack {
		res += num
	}
	return res
}
