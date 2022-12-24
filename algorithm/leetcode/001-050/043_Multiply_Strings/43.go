package _43_Multiply_Strings

import (
	"bytes"
	"fmt"
)

func multiply(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	max := l1 + l2
	result := make([]uint8, max, max)
	for i := l2 - 1; i >= 0; i-- {
		for j := l1 - 1; j >= 0; j-- {
			n2 := num2[i] - '0'
			n1 := num1[j] - '0'
			total := n1 * n2
			temp := result[i+j+1] + total%10
			result[i+j+1] = temp % 10
			result[i+j] = result[i+j] + total/10 + temp/10
		}
	}
	var res bytes.Buffer
	for i := range result {
		if res.Len() == 0 && result[i] == 0 {
			continue
		}
		res.WriteString(fmt.Sprintf("%d", result[i]))
	}
	if res.Len() == 0 {
		return "0"
	}
	return res.String()
}
