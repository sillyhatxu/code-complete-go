package main

func IsUniqueChars1(input string) bool {
	if len(input) > 128 {
		return false
	}
	check := make(map[int32]bool)
	for _, c := range input {
		_, ok := check[c]
		if ok {
			return false
		}
		check[c] = true
	}
	return true
}

func IsUniqueChars2(input string) bool {
	if len(input) > 128 {
		return false
	}
	check := make([]int32, 128)
	for _, c := range input {
		ok := check[c]
		if ok == 1 {
			return false
		}
		check[c] = 1
	}
	return true
}

//TODO 没看懂
func IsUniqueChars3(input string) bool {
	checker := 0
	for i, c := range input {
		val := c - 'a'
		if (checker & (i << val)) > 0 {
			return false
		}
		checker |= 1 << val
	}
	return true
}
