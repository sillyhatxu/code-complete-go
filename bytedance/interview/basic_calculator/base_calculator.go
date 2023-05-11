package basic_calculator

import (
	"strconv"
	"unicode"
)

func calculate(s string) int {
	// 定义一个运算符栈和一个操作数栈
	var signQueue []rune
	var numberQueue []int
	// 定义一个函数用于计算两个数的结果
	calc := func() {
		num1 := numberQueue[len(numberQueue)-1]
		num2 := numberQueue[len(numberQueue)-2]
		option := signQueue[len(signQueue)-1]
		numberQueue = numberQueue[:len(numberQueue)-2]
		signQueue = signQueue[:len(signQueue)-1]
		switch option {
		case '+':
			numberQueue = append(numberQueue, num2+num1)
		case '-':
			numberQueue = append(numberQueue, num2-num1)
		case '*':
			numberQueue = append(numberQueue, num2*num1)
		case '/':
			numberQueue = append(numberQueue, num2/num1)
		}
	}

	// 遍历字符串中的每个字符
	for i := 0; i < len(s); i++ {
		if unicode.IsSpace(rune(s[i])) {
			continue
		}
		curr := s[i]
		switch curr {
		case '+', '-':
			// 如果当前字符是 + 或 -，则计算栈中的操作数和运算符
			for len(signQueue) > 0 && (signQueue[len(signQueue)-1] == '+' || signQueue[len(signQueue)-1] == '-' || signQueue[len(signQueue)-1] == '*' || signQueue[len(signQueue)-1] == '/') {
				calc()
			}
			signQueue = append(signQueue, rune(curr))
		case '*', '/':
			// 如果当前字符是 * 或 /，则计算栈中的操作数和运算符
			for len(signQueue) > 0 && (signQueue[len(signQueue)-1] == '*' || signQueue[len(signQueue)-1] == '/') {
				calc()
			}
			signQueue = append(signQueue, rune(curr))
		case '(':
			// 如果当前字符是左括号，将其压入运算符栈
			signQueue = append(signQueue, rune(curr))
		case ')':
			// 如果当前字符是右括号，则计算栈中的操作数和运算符，直到遇到左括号
			for signQueue[len(signQueue)-1] != '(' {
				calc()
			}
			signQueue = signQueue[:len(signQueue)-1] // 弹出左括号
		default:
			currentNumberSrc := string(curr)
			// 如果当前字符是数字，判断这个数字的长度（234+456）
			for j := i + 1; j < len(s); j++ {
				ch := s[j]
				if unicode.IsDigit(rune(ch)) {
					currentNumberSrc = currentNumberSrc + string(ch)
					i++
				} else {
					break
				}
			}
			val, _ := strconv.Atoi(currentNumberSrc)
			numberQueue = append(numberQueue, val)
		}
	}
	// 计算栈中的操作数和运算符，直到运算符栈为空
	for len(signQueue) > 0 {
		calc()
	}
	// 最终操作数栈中剩余的元素即为计算结果
	return numberQueue[0]
}
