package _143_Longest_Common_Subsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	initialDP := func() [][]int {
		dp := make([][]int, len(text1)+1, len(text1)+1)
		dp[0] = make([]int, len(text2)+1)
		for i := 1; i <= len(text1); i++ {
			dp[i] = make([]int, len(text2)+1)
			for j := 1; j <= len(text2); j++ {
				dp[i][j] = -1
			}
		}
		return dp
	}
	dp := initialDP()
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = 1 + dp[i][j]
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
