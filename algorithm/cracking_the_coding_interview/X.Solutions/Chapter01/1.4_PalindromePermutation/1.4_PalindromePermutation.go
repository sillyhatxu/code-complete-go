package x_solutions_chapter01_PalindromePermutation

import "strings"

func PalindromePermutation(input string) bool {
	check := map[byte]int{}
	inputByte := []byte(strings.ToUpper(input))
	for i := 0; i < len(inputByte); i++ {
		if inputByte[i] == ' ' {
			continue
		}
		_, ok := check[inputByte[i]]
		if ok {
			check[inputByte[i]] = check[inputByte[i]] + 1
		} else {
			check[inputByte[i]] = 1
		}
	}
	oddCount := 0
	for _, v := range check {
		if v%2 == 1 {
			oddCount++
		}
		if oddCount > 1 {
			return false
		}
	}
	return true
}
