# 1312. Shortest Palindrome

A palindrome is defined as a word that reads the same forward as it does backward. For Example, The word "tacocat" is a palindrome, but the words "taco" and "cat" are not. Determine the minimum number of characters that must be inserted into a string to make it a palindrome.

For example, the string s = "abcda" can be made a palindrome by performing the following sequence of two character insertions: "abcda" -> "abdcda" -> "abdcdba".

## Function Description

Complete the function shortestPalindrome in the editor below. The function must return an integer that denotes the minimum number of characters that must be inserted into the string to make it a palindrome.

shortestPalindrome has the following parameter(s):
```
   s: a string
```

## Constraints

* 1 <= s <= 10^3
* String s consists of lowercase English alphabetic letters only, ascii[a-z]

## Input Format for Custom Testing

Input from stdin will be processed as follows and passed to the function.

The line contains the string s.

### Example 1:

```
Input: abcda
Output: 2
"abcda" -> "abdcda" -> "abdcdba"
```

### Example 2:

```
Input: abab
Output: 1
"abab" -> "ababa"
```

### Example 3:

```
Input: s = "zzazz"
Output: 0
Explanation: The string "zzazz" is already palindrome we do not need any insertions.
```

## Example 4:

```
Input: s = "mbadm"
Output: 2
Explanation: String can be "mbdadbm" or "mdbabdm".
```

## Example 5:

```
Input: s = "leetcode"
Output: 5
Explanation: Inserting 5 characters the string becomes "leetcodocteel".
```

### Translate:

> 1312. 让字符串成为回文串的最少插入次数

给你一个字符串s，每一次操作你都可以在字符串的任意位置插入任意字符。

请你返回让s成为回文串的最少操作次数。

**「回文串」是正读和反读都相同的字符串。**

## 提示：

* 1 <= s.length <= 500
* s 中所有字符都是小写字母。

## 解题思路

在这个解决方案中，我们首先定义了一个 minInsertions 函数，它将一个字符串作为输入，并返回将其转换为回文所需的最小插入次数。我们还定义了一个 min 函数，用于获取两个整数的最小值。

在 minInsertions 函数中，我们首先获取输入字符串的长度，并创建一个二维的 dp 数组，用于存储每个子串的最小插入次数。然后，我们使用动态规划算法，从后往前遍历字符串，计算出每个子串的最小插入次数。

对于任何两个字符串 s[i] 和 s[j]，如果它们相等，则它们的最小插入次数与子串 s[i+1:j-1] 的最小插入次数相同。否则，我们可以将 s[i] 插入到子串 s[i+1:j] 的末尾，或将 s[j] 插入到子串 s[i:j-1] 的开头，取两者中的较小值加 1。

最后，我们返回整个字符串的最小插入次数，即 dp[0][n-1]。

通过这个解决方案，我们可以在Golang中计算出将一个字符串变成回文字符串所需的最小插入次数。