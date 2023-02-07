package main

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	dick := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var res []string
	src := dick[digits[0]-'0']
	for j := 0; j < len(src); j++ {
		res = append(res, string(src[j]))
	}
	for i := 1; i < len(digits); i++ {
		src := dick[digits[i]-'0']
		size := len(res)
		for j := 0; j < size; j++ {
			for k := 0; k < len(src); k++ {
				res = append(res, res[j]+string(src[k]))
			}
		}
		res = res[size:]
	}
	return res
}

func letterCombinations1(digits string) []string {
	if digits == "" {
		return []string{}
	}
	dick := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var queue []string
	for i := 0; i < len(digits); i++ {
		src := dick[digits[i]-'0']
		for len(queue) == 0 || len(queue[0]) == i {
			curr := ""
			if len(queue) != 0 {
				curr = queue[0]
				queue = queue[1:]
			}
			for j := 0; j < len(src); j++ {
				queue = append(queue, curr+string(src[j]))
			}
		}
	}
	return queue
}

func switchLetter(letter byte) []string {
	switch {
	case letter == '2':
		return []string{"a", "b", "c"}
	case letter == '3':
		return []string{"d", "e", "f"}
	case letter == '4':
		return []string{"g", "h", "i"}
	case letter == '5':
		return []string{"j", "k", "l"}
	case letter == '6':
		return []string{"m", "n", "o"}
	case letter == '7':
		return []string{"p", "q", "r", "s"}
	case letter == '8':
		return []string{"t", "u", "v"}
	case letter == '9':
		return []string{"w", "x", "y", "z"}
	default:
		return []string{}
	}
}

func merge(a1, a2 []string) []string {
	var result []string
	for _, v1 := range a1 {
		for _, v2 := range a2 {
			result = append(result, v1+v2)
		}
	}
	return result
}

func letterCombinationsOriginal(digits string) []string {
	if digits == "" {
		return []string{}
	} else if len(digits) == 1 {
		return switchLetter(digits[0])
	}
	result := switchLetter(digits[0])
	for i := 1; i < len(digits); i++ {
		result = merge(result, switchLetter(digits[i]))
	}
	return result
}
