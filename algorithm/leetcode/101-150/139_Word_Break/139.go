package _39_Word_Break

import "fmt"

//TODO 待查看
//func wordBreak(s string, wordDict []string) bool {
//	dp := make([]bool, len(s)+1)
//	dp[len(s)] = true
//	for i := len(s) - 1; i >= 0; i-- {
//		for _, word := range wordDict {
//			if i+len(word) <= len(s) && s[i:i+len(word)] == word {
//				dp[i] = dp[i+len(word)]
//			}
//			if dp[i] {
//				break
//			}
//		}
//	}
//	return dp[0]
//}

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			temp := s[j:i]
			fmt.Println(temp)
			if _, ok := wordMap[temp]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
