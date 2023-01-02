# 763. Partition Labels

You are given a string s. We want to partition the string into as many parts as possible so that each letter appears in at most one part.

Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.

Return a list of integers representing the size of these parts.

### Example 1:

```
Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.
```

### Example 2:

```
Input: s = "eccbbbbdec"
Output: [10]
```

### Constraints:

* 1 <= s.length <= 500
* s consists of lowercase English letters.

### Translate:

> 763. 划分字母区间

给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。

注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。

返回一个表示每个字符串片段的长度的列表。

### 示例 1：

```
输入：s = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca"、"defegde"、"hijhklij" 。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 这样的划分是错误的，因为划分的片段数较少。
```

### 示例 2：

```
输入：s = "eccbbbbdec"
输出：[10]
```

### 提示：

* 1 <= s.length <= 500
* s 仅由小写英文字母组成


# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
<!-- Describe your approach to solving the problem. -->
```
    ababcbaca defegdehijhklij
start 0; end:8
    for loop [ababcbaca] find last index is 8
        res [8-0+1]; res:[9]; start:9; end:8
    for loop [defegd] find last index is 14 (ehijhklij)
        but foloop i = 10; the last index is 15; So start:9; end:15
    for loop [defegde] find last index is 15 (hijhklij)
        res [15-9+1]; res:[9,7]; start:16; end:15
    for loop [hijh] find last index is 19 (klij)
        but alphabet [i] the last index is 22
        and alphabet [j] the last index is 23
        res [23-16+1]; res:[9,7,8]
```

# Complexity
- Time complexity: O(n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: O(1)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
func partitionLabels(s string) []int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	lastIndex := make([]int, 26, 26)
	for i := 0; i < len(s); i++ {
		lastIndex[s[i]-'a'] = i
	}
	var res []int
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		end = max(end, lastIndex[s[i]-'a'])
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}
```