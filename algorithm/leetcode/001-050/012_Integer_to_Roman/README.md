# 12. Integer to Roman

Roman numerals are represented by seven different symbols: `I`, `V`, `X`, `L`, `C`, `D` and `M`.

```
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```

For example, two is written as `II` in Roman numeral, just two one's added together. Twelve is written as, `XII`, which is simply `X` + `II`. The number twenty seven is written as `XXVII`, which is `XX` + `V` + `II`.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not `IIII`. Instead, the number four is written as `IV`. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as `IX`* . There are six instances where subtraction is used:
-
* `I` can be placed before `V` (5) and `X` (10) to make 4 and 9. 
* `X` can be placed before `L` (50) and `C` (100) to make 40 and 90. 
* `C` can be placed before `D` (500) and `M` (1000) to make 400 and 900.

Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

### Example 1:

```
Input: 3
Output: "III"
```

### Example 2:

```
Input: 4
Output: "IV"
```

### Example 3:

```
Input: 9
Output: "IX"
```

### Example 4:

```
Input: 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
```

### Example 5:

```
Input: 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```

### Translate:

> 12. 整数转罗马数字

罗马数字包含以下七种字符：I，V，X，L，C，D和M。
```
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```
例如， 罗马数字 2 写做II，即为两个并列的 1。12 写做XII，即为X+II。 27 写做XXVII, 即为XX+V+II。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做IIII，而是IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为IX。这个特殊的规则只适用于以下六种情况：

* I可以放在V(5) 和X(10) 的左边，来表示 4 和 9。
* X可以放在L(50) 和C(100) 的左边，来表示 40 和90。
* C可以放在D(500) 和M(1000) 的左边，来表示400 和900。

给你一个整数，将其转为罗马数字。

### 示例1:

```
输入:num = 3
输出: "III"
```

### 示例2:

```
输入:num = 4
输出: "IV"
```

### 示例3:

```
输入:num = 9
输出: "IX"
```

### 示例4:

```
输入:num = 58
输出: "LVIII"
解释: L = 50, V = 5, III = 3.
```

### 示例5:

```
输入:num = 1994
输出: "MCMXCIV"
解释: M = 1000, CM = 900, XC = 90, IV = 4.
```