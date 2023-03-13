package interview

/**
在这个解决方案中，我们首先定义了一个 minInsertions 函数，它将一个字符串作为输入，并返回将其转换为回文所需的最小插入次数。我们还定义了一个 min 函数，用于获取两个整数的最小值。

在 minInsertions 函数中，我们首先获取输入字符串的长度，并创建一个二维的 dp 数组，用于存储每个子串的最小插入次数。然后，我们使用动态规划算法，从后往前遍历字符串，计算出每个子串的最小插入次数。

对于任何两个字符串 s[i] 和 s[j]，如果它们相等，则它们的最小插入次数与子串 s[i+1:j-1] 的最小插入次数相同。否则，我们可以将 s[i] 插入到子串 s[i+1:j] 的末尾，或将 s[j] 插入到子串 s[i:j-1] 的开头，取两者中的较小值加 1。

最后，我们返回整个字符串的最小插入次数，即 dp[0][n-1]。

通过这个解决方案，我们可以在Golang中计算出将一个字符串变成回文字符串所需的最小插入次数。
*/
func minInsertions(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
