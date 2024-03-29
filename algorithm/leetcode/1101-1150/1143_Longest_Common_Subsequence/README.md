# 1143. Longest Common Subsequence

Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

* For example, "ace" is a subsequence of "abcde".

A common subsequence of two strings is a subsequence that is common to both strings.


### Example 1:

```
Input: text1 = "abcde", text2 = "ace" 
Output: 3  
Explanation: The longest common subsequence is "ace" and its length is 3.
```

### Example 2:

```
Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.
```

### Example 3:

```
Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.
```

### Constraints:

* 1 <= text1.length, text2.length <= 1000
* text1 and text2 consist of only lowercase English characters.

### Translate:

> 1143. 最长公共子序列

给定两个字符串text1 和text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

* 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。

两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

### 示例 1：

```
输入：text1 = "abcde", text2 = "ace"
输出：3  
解释：最长公共子序列是 "ace" ，它的长度为 3 。
```

### 示例 2：

```
输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc" ，它的长度为 3 。
```
### 示例 3：

```
输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0 。
```

### 提示：

* 1 <= text1.length, text2.length <= 1000
* text1 和text2 仅由小写英文字符组成。

# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity: O(m*n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: O(m*n)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
func longestCommonSubsequence(text1 string, text2 string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	initialDP := func() [][]int{
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
```