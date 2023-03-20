# 91. Decode Ways

A message containing letters from A-Z can be encoded into numbers using the following mapping:

```
'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
```

To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above (there may be multiple ways). For example, "11106" can be mapped into:

* "AAJF" with the grouping (1 1 10 6)
* "KJF" with the grouping (11 10 6)

Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".

Given a string s containing only digits, return the number of ways to decode it.

The test cases are generated so that the answer fits in a 32-bit integer.



### Example 1:

```
Input: s = "12"
Output: 2
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
```

### Example 2:

```
Input: s = "226"
Output: 3
Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
```

### Example 3:

```
Input: s = "06"
Output: 0
Explanation: "06" cannot be mapped to "F" because of the leading zero ("6" is different from "06").1 <= s.length <= 100
```

### Constraints:

* 1 <= s.length <= 100
* s contains only digits and may contain leading zero(s).

### Translate:

> 91. 解码方法

一条包含字母A-Z 的消息通过以下映射进行了 编码 ：

'A' -> "1"
'B' -> "2"
...
'Z' -> "26"

要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：

* "AAJF" ，将消息分组为 (1 1 10 6)
* "KJF" ，将消息分组为 (11 10 6)

注意，消息不能分组为 (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。

给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。

题目数据保证答案肯定是一个 32 位 的整数。

示例 1：

```
输入：s = "12"
输出：2
解释：它可以解码为 "AB"（1 2）或者 "L"（12）。
```

示例 2：

```
输入：s = "226"
输出：3
解释：它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
```

示例 3：

```
输入：s = "0"
输出：0
解释：没有字符映射到以 0 开头的数字。
含有 0 的有效映射是 'J' -> "10" 和 'T'-> "20" 。
由于没有字符，因此没有有效的方法对此进行解码，因为所有数字都需要映射。
```

提示：

1 <= s.length <= 100
s 只包含数字，并且可能包含前导零。


### 解题思路

DP 0ms 100%; Time: O(n), Space: O(1)

首先在不考虑0的前提下

```
DP[1] = 1
DP[2] = DP[1] + DP[1] = 2
DP[3] = DP[2] + DP[1] = 3
...
DP[i] = DP[i-1] + DP[i-2] = ?
```

### 举例：

输入: `12101226`

默认previous和current也就是dp[i-1], dp[i-2] 为1，因为此时i的值为0

所以遍历时直接从i=1开始遍历

#### i = 1
```
previous=1
current=1
  i
1 2 1 0 1 2 2 6
___

10 <= 12 <= 26,并且当前索引位置不是0，所以DP[2] = DP[1] + DP[1] = 2,所以 previous赋值现在current的值，current = current + previous
```

#### i = 2

```
previous=1
current=2
    i
1 2 1 0 1 2 2 6
  ___

10 <= 21 <= 26,并且当前索引位置不是0，所以DP[3] = DP[2] + DP[1] = 3,所以 previous赋值现在current的值，current = current + previous
```

#### i = 3

```
previous=2
current=3
      i
1 2 1 0 1 2 2 6
    ___

10 <= 10 <= 26,但是当前索引位置是0，无效; 所以DP[4] = DP[3] + 0 = 3 + 0 = 3,所以 current = previous = 2,previous = x(随便给值就行)

```

#### i = 4

```
previous=0
current=2
        i
1 2 1 0 1 2 2 6
      ___

0 < 1 <= 9,并且当前索引位置不是0，DP结果保持不变，赋值给previous； previous = current
```

#### i = 5

```
previous=2
current=2
          i
1 2 1 0 1 2 2 6
        ___

10 <= 14 <= 26,并且当前索引位置不是0，所以 previous赋值现在current的值，current = current + previous
```

#### i = 6

```
previous=2
current=4
            i
1 2 1 0 1 2 2 6
          ___

10 <= 22 <= 26,并且当前索引位置不是0，所以 previous赋值现在current的值，current = current + previous
```

#### i = 7

```
previous=4
current=6
              i
1 2 1 0 1 2 2 6
            ___

10 <= 26 <= 26,并且当前索引位置不是0，所以 previous赋值现在current的值，current = current + previous = 10
```

最后结果为10


`getRealNum`意思是：字节'0' ~ '9', golang中的数字为从 48 ~ 57(其实不需要管实际数字)

在不是用toString的前提下，我们可以直接拿到 9 (byte)

```
'9' - '0' = 57 - 48 = 9
```

### 代码

```golang
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	//dp[i] = dp[i-1] + dp[i-2]
	current, previous := 1, 1 //dp[i], dp[i-1]
	for i := 1; i < len(s); i++ {
		num := getRealNum(s[i-1])*10 + getRealNum(s[i])
		if num >= 10 && num <= 26 {
			//[10-26]
			if s[i] == '0' {
				current, previous = previous, 0
			} else {
				previous, current = current, current+previous
			}
		} else {
			//(0-9]
			if s[i] == '0' {
				return 0
			} else {
				previous = current
			}
		}
	}
	return current
}

func getRealNum(input byte) byte {
	return input - '0'
}
```


### index

```
数字转字母，几种方式
```