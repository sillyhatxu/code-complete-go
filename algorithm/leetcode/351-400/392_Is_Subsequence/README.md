# 392. Is Subsequence

Given two strings `s` and `t`, return `true` if `s` is `a` subsequence of `t`, or `false` otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "`ace`" is a subsequence of "`abcde`" while "`aec`" is not).


### Example 1:


```
Input: s = "abc", t = "ahbgdc"
Output: true
```

### Example 2:


```
Input: s = "axc", t = "ahbgdc"
Output: false
```


### Constraints:

* 0 <= s.length <= 100
* 0 <= t.length <= 10^4
* `s` and `t` consist only of lowercase English letters.

### Translate:

> 392. 判断子序列
> 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
> 
> 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
> 
> 进阶：
> 
> 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
> 
> 提示：
> 
> 0 <= s.length <= 100
> 
> 0 <= t.length <= 10^4
> 
> 两个字符串都只由小写字符组成。