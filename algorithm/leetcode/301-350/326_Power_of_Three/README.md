# 326. Power of Three

Given an integer n, return true if it is a power of three. Otherwise, return false.

An integer n is a power of three, if there exists an integer x such that n == 3x.

### Example 1:

```
Input: n = 27
Output: true
Explanation: 27 = 33
```

### Example 2:

```
Input: n = 0
Output: false
Explanation: There is no x where 3x = 0.
```

### Example 3:

```
Input: n = -1
Output: false
Explanation: There is no x where 3x = (-1).
```

### Constraints:

* -231 <= n <= 231 - 1

### Follow up

Could you solve it without loops/recursion?

### Translate:

> 326. 3 的幂

给定一个整数，写一个函数来判断它是否是 3的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x

进阶：你能不使用循环或者递归来完成本题吗？


### 解题思路

最终解决方案

### non-loops

```golang
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	quo,_ := math.Modf(math.Log10(float64(n)) / math.Log10(3))
	return n == int(math.Pow(3, quo))
}
```

### loops

```golang
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}
```

已知下边的公式
![image description](WechatIMG415.jpeg)

所以我们只需要判断x是否为整数

```golang
func isPowerOfThree(n int) bool {
	return (math.Log10(float64(n)) / math.Log10(float64(3))) % 1 == 0
}
````

但是golang上边的代码会报编译错误 **(operator % is not defined on untyped int)**

所以我们需要换一种写法
```golang
func isPowerOfThree(n int) bool {
    temp := math.Log10(float64(n)) / math.Log10(float64(3))
    _, res := math.Modf(temp)
    return res == 0.0
}
```

但是这种会有精度问题

> temp := math.Log10(float64(27)) / math.Log10(float64(3))
> 
> temp = 3.0000000000000004