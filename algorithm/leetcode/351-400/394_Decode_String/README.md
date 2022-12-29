# 394. Decode String

Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 10^5.


### Example 1:

```
Input: s = "3[a]2[bc]"
Output: "aaabcbc"
```

### Example 2:

```
Input: s = "3[a2[c]]"
Output: "accaccacc"
```

### Example 3:

```
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"
```

### Constraints:

* 1 <= s.length <= 30
* s consists of lowercase English letters, digits, and square brackets '[]'.
* s is guaranteed to be a valid input.
* All the integers in s are in the range [1, 300].

### Translate:

> 394. 字符串解码

给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像3a或2[4]的输入。

### 提示：

* 1 <= s.length <= 30
* s由小写英文字母、数字和方括号'[]' 组成
* s保证是一个有效的输入。
* s中所有整数的取值范围为[1, 300]


# Best solution; recursion and double stack;

# Recursion Solution

### Complexity
- Time complexity: O(n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: O(1)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

### Code
```
func decodeString(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			res += s[i : i+1]
		} else {
			temp, num, c := "", 0, s[i]-'0'
			for c <= 9 {
				num = num*10 + int(c)
				i++
				c = s[i] - '0'
			}
			if s[i] == '[' {
				index, curr := i, 0
				for index < len(s) {
					if s[index] == '[' {
						curr++
					}
					if s[index] == ']' {
						curr--
					}
					if curr == 0 {
						break
					}
					index++
				}
				temp = decodeString(s[i+1 : index])
				i = index
			}
			for j := 0; j < num; j++ {
				res += temp
			}
		}
	}
	return res
}
```

# Double Stack Solution

### Complexity
- Time complexity: O(n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: O(n)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

### Code
```
func decodeString(s string) string {
	var nums []int
	var dicks []string
	for i := 0; i < len(s); i++ {
		c := s[i] - '0'
		if c <= 9 {
			//number
			num := int(c)
			for {
				next := s[i+1] - '0'
				if next > 9 {
					break
				}
				num = num*10 + int(next)
				i++
			}
			nums = append(nums, num)
		} else if s[i] == ']' {
			index, temp := len(dicks)-1, ""
			for dicks[index] != "[" {
				temp = dicks[index] + temp
				index--
			}
			dicks = dicks[:index]
			curr, times := "", nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			for j := 0; j < times; j++ {
				curr += temp
			}
			dicks = append(dicks, curr)
		} else {
			//letter or '['
			dicks = append(dicks, string(s[i]))
		}
	}
	var res bytes.Buffer
	for i := 0; i < len(dicks); i++ {
		res.WriteString(dicks[i])
	}
	return res.String()
}
```