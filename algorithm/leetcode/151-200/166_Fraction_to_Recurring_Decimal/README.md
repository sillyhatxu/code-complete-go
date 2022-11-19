# 166. Fraction to Recurring Decimal

Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.

If the fractional part is repeating, enclose the repeating part in parentheses.

If multiple answers are possible, return any of them.

It is guaranteed that the length of the answer string is less than 104 for all the given inputs.

### Example 1:

```
Input: numerator = 1, denominator = 2
Output: "0.5"
```

### Example 2:

```
Input: numerator = 2, denominator = 1
Output: "2"
```

### Example 3:

```
Input: numerator = 4, denominator = 333
Output: "0.(012)"
```

### Constraints:

* -231 <= numerator, denominator <= 231 - 1
* denominator != 0

### Translate:

> 166. 分数到小数

给定两个整数，分别表示分数的分子numerator 和分母 denominator，以 字符串形式返回小数 。

如果小数部分为循环小数，则将循环的部分括在括号内。

如果存在多个答案，只需返回 任意一个 。

对于所有给定的输入，保证 答案字符串的长度小于 104 。