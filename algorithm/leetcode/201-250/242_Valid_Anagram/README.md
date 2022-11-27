# 242. Valid Anagram

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

### Example 1:

```
Input: s = "anagram", t = "nagaram"
Output: true
```

### Example 2:

```
Input: s = "rat", t = "car"
Output: false
```

### Constraints:

* 1 <= s.length, t.length <= 5 * 104
* s and t consist of lowercase English letters.

### Follow up

What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

### Translate:

> 242. 有效的字母异位词

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。

提示:

* 1 <= s.length, t.length <= 5 * 104
* s 和 t仅包含小写字母

进阶:如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？