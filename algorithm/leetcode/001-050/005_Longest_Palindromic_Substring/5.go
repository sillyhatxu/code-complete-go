package main

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
			curr := s[j : i+1]
			if isPalindrome(curr) {
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
