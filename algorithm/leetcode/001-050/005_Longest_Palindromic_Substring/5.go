package main

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 || (len(s) == 2 && s[0] == s[1]) {
		//s = ""; s = "a"; s = "aa"
		return s
	}
	isPalindrome := func(i string) bool {
		l, r := 0, len(i)-1
		for l < r {
			if i[l] != i[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	pre, res := s[:1], s[:1]
	for right := 1; right < len(s); right++ {
		left := right - len(pre) - 1
		if left < 0 {
			left = 0
		}
		for j := left; j <= right; j++ {
			if isPalindrome(s[j : right+1]) {
				pre = s[j : right+1]
				if len(pre) > len(res) {
					res = pre
				}
				break
			}
		}
	}
	return res
}

//最简单的理解方式，不过很现实时间复杂度太高了，LeetCode已经超时
func longestPalindromeLow(s string) string {
	if len(s) == 1 || (len(s) == 2 && s[0] == s[1]) {
		return s
	}
	var sonSrc []string
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			sonSrc = append(sonSrc, s[i:j])
		}
	}
	maxSrc := ""
	for _, v := range sonSrc {
		negate := ""
		for i := len(v) - 1; i >= 0; i-- {
			negate += string(v[i])
		}
		if v == negate && len(v) > len(maxSrc) {
			maxSrc = v
		}
	}
	return maxSrc
}

//Dynamic Programming
//使用动态规划来解决
//执行结果
/**
Runtime: 64 ms, faster than 38.83% of Go online submissions for Longest Palindromic Substring.
Memory Usage: 6.5 MB, less than 18.26% of Go online submissions for Longest Palindromic Substring.
Next challenges:
*/
func longestPalindromeOptimize(s string) string {
	if len(s) == 0 || len(s) == 1 || (len(s) == 2 && s[0] == s[1]) {
		return s
	}
	dynamicProgrammingArray := make([][]bool, len(s))
	left, right, length := 0, 0, 0
	for i := 0; i < len(s); i++ {
		temp := make([]bool, len(s))
		dynamicProgrammingArray[i] = temp
		dynamicProgrammingArray[i][i] = true
		for j := 0; j <= i; j++ {
			dynamicProgrammingArray[j][i] = s[i] == s[j] && (i-j < 2 || dynamicProgrammingArray[j+1][i-1])
			if dynamicProgrammingArray[j][i] && length < i-j+1 {
				length = i - j + 1
				left = j
				right = i
			}
		}
	}
	return s[left : right+1]
}

/**
LeetCode 中最牛逼解法

运行结果
Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Palindromic Substring.
Memory Usage: 2.2 MB, less than 98.26% of Go online submissions for Longest Palindromic Substring.
*/
func longestPalindromeBest(s string) string {
	if s == "" {
		return ""
	}
	pre, max := s[:1], s[:1]
	for i := 1; i < len(s); i++ {
		begin := i - len(pre) - 1
		if begin < 0 {
			begin = 0
		}
		for j := begin; j <= i; j++ {
			if isPalindrome(s[j : i+1]) {
				pre = s[j : i+1]
				if len(pre) > len(max) {
					max = pre
				}
				break
			}
		}
	}
	return max
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
