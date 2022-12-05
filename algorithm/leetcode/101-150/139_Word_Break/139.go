package _39_Word_Break

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = struct{}{}
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			temp := s[j:i]
			if _, ok := wordMap[temp]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
