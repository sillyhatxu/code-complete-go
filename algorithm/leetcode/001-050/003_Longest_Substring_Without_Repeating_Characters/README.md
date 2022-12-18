# 3.Longest Substring Without Repeating Characters

Given a string s, find the length of the longest substring without repeating characters.

### Example 1:

```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

### Example 2:

```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

### Example 3:

```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

### Constraints:

* 0 <= s.length <= 5 * 104
* s consists of English letters, digits, symbols and spaces.

### Translate:

> 3. 无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。

### 示例1:

```
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

### 示例 2:

```
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
```

### 示例 3:

```
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
```

### 提示：

* 0 <= s.length <= 5 * 10^4
* s由英文字母、数字、符号和空格组成

#### 解析

![](https://lh3.googleusercontent.com/xFYPKT9XTgxFBFYB3P3YERU5wgOAwTZP20f3zXzMxWWXRbo7gd2Gi_8FuQXVtYvDDYK6Wgss-n9L__zq5KMvn3OG0gwApp1IeJe_d84wk3eRWLFaerAFHCuiiI1sjP3IAvdBUd9Og9cX02uNPbGd6Iksb2PedTNw2X51h_ze1TNvWmclRdcLf7mrwjkxX96KnVdFGtesLpjVcR2_GB7d6nQpc4WsCUsMwrPo-I_nBvGG-A5izNc0jAAAmhhmFpyzJqXai8Ao6MKCI77p2tf-FgT9_o3jVQrT7LI06xih0JQLv83V5dFSS2bnxeSpaogeCBrL1GpBIObw31sBmxdQzlLD5GiviXjrQo6H6LF59BpvNIdUYjWTG1X9Gwy5FrWXFwTTB5lokNFQLTPX11SYoBFBWFhvEq8ykllmdxyxo_wAxUxGhaVrFCFsDdHAlHhXcYgkhe2a1mTGnniEOTpq0b7ykCISIzx0Mm-vw6GEiYL9FJE-c3ml0w0-N-nx_aZyqiLpjzjsa6NCkq7cVYSbYQUdlW9TURD9nxBrKFWWkM0-VNsJJJsJmJzqH1S_5-M1qiNYkMV_dUk644okFj1iNKziVgkV6o3cgeXzaB_PLr3T7LhD4TkNhtgUw0ZbKBnnpSZW0pafsrnXLehHSLm3XCKygxwa2Oo=w852-h479-no)

---

![](https://lh3.googleusercontent.com/_cT__EOgHahpZXk0yrcjkUXc-IAO7eMiAX8nJHudULFaHUFOWZ7lzC4PoZpltdqf0rZ1bkWaeiHx4RBrue10vZJK0CkT7WSJEGvN1R7vGtWKt58cNHsdUuGCX9orCNnCQ-Z-grXjOMC9xGkNeVvFkkaDVlUa965_iRmI2Bq9CNydEHftkDuzIDTP6HsS2cSajmGo4yoJnWNHyNRiqOXerBpPLjmRwVX0VO2iu29kUpoGCqZg0fok3KL1VLnPovAgDrqZ17jlJKJi-hMLj1Wnw8w_RVv5uMMmU9qZVhTWGv9BR0ulKdTNmllzDtvUEYphroWQlOFy-XlfD6geKpxoVlpNzorcz_HX9X5cD6EU7-P-MSEKwnEX55RB1GOoFKOT53NSU_0x6sS9Jsy4HIdwNnacr_do5iLe5L34B1z65lzRIxTTnvGl65jkp6qIHaSFCRWmxldEAXYLMQktRFg1NczC0dvI-qSlxL7C1uAjFB_uq-CVA6usUoJkOqYDs2R0_vD5zwamwBSrixxD6UVnOybsSMBYEuyUr-9KaRB3Gaj999x-oHw7KdWIdtLMRnSX-eTCuG1P4yCSChp_XvYQvSyvnu1eHNOF1etjazg3j7Xb1r_uDx4R_dVh062SgcULXp7olb_0YztMeTOpNZPoKHKzlUsKyww=w745-h897-no)
