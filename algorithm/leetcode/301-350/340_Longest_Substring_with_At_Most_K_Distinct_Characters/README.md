# 340. Longest Substring with At Most K Distinct Characters

Given a string, find the length of the longest substring T that contains at most k distinct characters.

### Example 1:

```
Input: s = "eceba", k = 2
Output: 3
Explanation: T is "ece" which its length is 3.
```

### Example 2:

```
Input: s = "aa", k = 1
Output: 2
Explanation: T is "aa" which its length is 2.
```

This question is a typical question that can be solved by using our substring template. That is, we keep the status of the substring into a HashMap, based on the specification of the question. This question requires us to find the longest substring in which contains at most k duplicate characters. Thus we keep a HashMap to store the characters and their number we have encountered in the current substring.

The substring is represented by two pointers labelled as start and end. Once we have less than or equal to k characters in our sliding window, we update the max length recorded is necessary and move the end pointer forward. Once we have more than k characters in our sliding window, we move the start pointer forward until we remove one character out from our HashMap (indicated by the count of the character turns into zero).

Note that the HashMap in this solution can be substituted using a nums[256] as well. In this case we need to keep an extra counter for the number of characters rather than calling the HashMap.size() function.

Time complexity: O(n) time where n is the number of characters in the input string.

Space complexity: O(n) space in the worst case when k is greater than s.length() and s contains only different characters in it.

# Translate: 340. 至多包含 K 个不同字符的最长子串

给你一个字符串 s 和一个整数 k ，请你找出 至多 包含 k 个 不同 字符的最长子串，并返回该子串的长度。

### Example 1:

```
输入：s = "eceba", k = 2
输出：3
解释：满足题目要求的子串是 "ece" ，长度为 3 。
```

### Example 2:

```
输入：s = "aa", k = 1
输出：2
解释：满足题目要求的子串是 "aa" ，长度为 2 。
```

### Constraints:

* 1 <= s.length <= 5 * 104
* 0 <= k <= 50