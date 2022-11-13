package _31_Palindrome_Partitioning

func partition(s string) [][]string {
	var res [][]string
	helper(&res, s, 0, []string{})
	return res
}

func helper(res *[][]string, s string, start int, temp []string) {
	if start == len(s) {
		copyTemp := make([]string, len(temp))
		copy(copyTemp, temp)
		*res = append(*res, copyTemp)
		return
	}
	for i := start + 1; i <= len(s); i++ {
		cur := s[start:i]
		if isPalindrome(cur) {
			temp = append(temp, cur)
			helper(res, s, i, temp)
			temp = temp[:len(temp)-1]
		}
	}
}

func isPalindrome(input string) bool {
	l, r := 0, len(input)-1
	for l < r {
		if input[l] != input[r] {
			return false
		}
		l++
		r--
	}
	return true
}
