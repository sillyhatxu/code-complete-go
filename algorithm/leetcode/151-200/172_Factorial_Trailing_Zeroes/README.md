# 172. Factorial Trailing Zeroes

Given an integer n, return the number of trailing zeroes in n!.

Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.

### Example 1:

```
Given an integer n, return the number of trailing zeroes in n!.

Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.
```

### Example 2:

```
Input: n = 5
Output: 1
Explanation: 5! = 120, one trailing zero.
```

### Example 3:

```
Input: n = 0
Output: 0
```

### Constraints:

* 0 <= n <= 104

### Translate:

> # 172. 阶乘后的零

给定一个整数 n ，返回 n! 结果中尾随零的数量。

提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1


### Example 1:

```
输入：n = 3
输出：0
解释：3! = 6 ，不含尾随 0
```

### Example 2:

```
输入：n = 5
输出：1
解释：5! = 120 ，有一个尾随 0
```

### Example 3:

```
输入：n = 0
输出：0
```

### Constraints:

* 0 <= n <= 104

### 进阶：

你可以设计并实现对数时间复杂度的算法来解决此问题吗？


### 解题思路

解题步骤分为两步
第一步，证明解题思路
第二部，证明

#### 第一步

这道题要找到n的阶层的得数Y，它的最后几位0的个数

也就是不管n是多少，假设Y=abc 000000 = abc * 10^6 = abc * (2*5) ^ 6

我们用10来做例子，将10!拆分成职数的乘积


```
10！
10      9      8       7       6       5       4       3       2       1
2*5     3*3    2*2*2   7      2*3      5       2*2     3       2       1
```

我们知道，要想得到这个数字最后等于一个整数(abc)*10^x 首先要有2和5，成对数出现，而2出现的次数，远远多于5出现的次数。

所以这道题根本上是数一数，n!中，可以拆分出5的个数。所以10！最后有2个0，答案为2

#### 第二步

### 代码

```golang
func trailingZeroes(n int) int {
    numOf5 := 0
	for n > 4 {
		numOf5, n = numOf5+n/5, n/5
	}
	return numOf5
}
```