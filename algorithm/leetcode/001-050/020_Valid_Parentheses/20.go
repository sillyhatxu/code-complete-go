package main

func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok {
			if len(stack) == 0 || stack[0] != v {
				return false
			}
			stack = stack[1:]
		} else {
			stack = append([]byte{s[i]}, stack...)
		}
	}
	return len(stack) == 0
}

/**
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Valid Parentheses.
Next challenges:
*/
func isValidOriginal(s string) bool {
	signMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	index := -1
	var stack = make([]string, len(s))
	for _, v := range s {
		if index == -1 || signMap[string(v)] != stack[index] {
			index++
			stack[index] = string(v)
		} else {
			index--
		}
	}
	return index == -1
}
