package main

import (
	"fmt"
	"strconv"
)

func Calculator(input string) int {
	// 定义一个运算符栈和一个操作数栈
	signQueue := []rune{}
	numberQueue := []int{}

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
	for _, curr := range input {
		switch curr {
		case '+', '-':
			// 如果当前字符是 + 或 -，则计算栈中的操作数和运算符
			for len(signQueue) > 0 && (signQueue[len(signQueue)-1] == '+' || signQueue[len(signQueue)-1] == '-' || signQueue[len(signQueue)-1] == '*' || signQueue[len(signQueue)-1] == '/') {
				calc()
			}
			signQueue = append(signQueue, curr)
		case '*', '/':
			// 如果当前字符是 * 或 /，则计算栈中的操作数和运算符
			for len(signQueue) > 0 && (signQueue[len(signQueue)-1] == '*' || signQueue[len(signQueue)-1] == '/') {
				calc()
			}
			signQueue = append(signQueue, curr)
		case '(':
			// 如果当前字符是左括号，将其压入运算符栈
			signQueue = append(signQueue, curr)
		case ')':
			// 如果当前字符是右括号，则计算栈中的操作数和运算符，直到遇到左括号
			for signQueue[len(signQueue)-1] != '(' {
				calc()
			}
			signQueue = signQueue[:len(signQueue)-1] // 弹出左括号
		default:
			// 如果当前字符是数字，则将其解析为整数，并将其压入操作数栈
			val, _ := strconv.Atoi(string(curr))
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

func main() {
	var res int

	test1 := "1+2+3"
	res = Calculator(test1)
	fmt.Printf("test1: %s, res: %v", test1, res)
	fmt.Println()
	test2 := "1+2*3"
	res = Calculator(test2)
	fmt.Printf("test1: %s, res: %v", test2, res)

}
