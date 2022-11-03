# 50. Pow(x, n)

Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

### Example 1:

```
Input: x = 2.00000, n = 10
Output: 1024.00000
```

### Example 2:

```
Input: x = 2.10000, n = 3
Output: 9.26100
```

### Example 3:

```
Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
```

### Constraints:

* -100.0 < x < 100.0
* -231 <= n <= 231-1
* n is an integer.
* -104 <= xn <= 104

### Translate:

> 50. Pow(x, n)
> 
> 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，x^n ）。


### 解题思路
首先，暴力写法直接for循环，Time complexity: O(n),所以肯定需要找一个更好的
可以由公式

```
X^n = X^2/n * X^2/n = X^4/n * X^4/n * X^4/n * X^4/n = ......
```

因此得出可以用类似于二分法解答。
Time complexity: O(logn)
Space complexity: O(logn)

### 代码

```golang
func myPow(x float64, n int) float64 {
	if n == 0 || x == 1{
		return 1
	}
	if n < 0 {
		return 1/recursion(x,n*-1)
	}
	return recursion(x,n)
}

func recursion(x float64, n int) float64{
    if n == 1{
        return x
    }
    half := recursion(x,n/2)
    if n % 2 == 0{
        return half * half
    }else{
        return half * half * x
    }
}
```