# 118. Pascal's Triangle

Given an integer `numRows`, return the first numRows of **Pascal's triangle**.

In **Pascal's triangle**, each number is the sum of the two numbers directly above it as shown:


### Example 1:

![image description](PascalTriangleAnimated2.gif)

```
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
```

### Example 2:

```
Input: numRows = 1
Output: [[1]]
```

### Constraints:

* 1 <= numRows <= 30

### Translate:

> 118. 杨辉三角

给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。


### 解题思路
```
                    (1[0][0])
                (1[1][0]) (1[1][1])
            (1[2][0]) (3[2][1]) (1[2][2])
        (1[3][0]) (4[3][1]) (4[3][2])  (1[3][3])
        
由图可知
1. 每行第一和最后一个元素为1
2. DP[2][1] = DP[1][0] + DP[1][1]
3. DP[3][1] = DP[2][0] + DP[2][1]
4. DP[3][2] = DP[2][1] + DP[2][2]
由此推出公式
DP[i][j] = DP[i-1][j-1] + DP[i-1][j]
```

### 代码

```golang
func generate(numRows int) [][]int {
	var result [][]int
	for i := 0; i < numRows; i++ {
		var levels []int
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				levels = append(levels, 1)
			} else {
				levels = append(levels, result[i-1][j] + result[i-1][j-1])
			}
		}
		result = append(result, levels)
	}
	return result
}
```