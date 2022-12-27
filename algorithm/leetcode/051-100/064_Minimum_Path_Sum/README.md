# 64. Minimum Path Sum

Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

**Note:** You can only move either down or right at any point in time.

### Example 1:

![image description](minpath.jpeg)

```
Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
```

### Example 2:

```
Input: grid = [[1,2,3],[4,5,6]]
Output: 12
```

### Constraints:

* m == grid.length
* n == grid[i].length
* 1 <= m, n <= 200
* 0 <= grid[i][j] <= 100

### Translate:

> 64. 最小路径和

给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

**说明：** 每次只能向下或者向右移动一步。

### 示例 1：

![image description](minpath.jpeg)

```
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
```

### 示例 2：

```
输入：grid = [[1,2,3],[4,5,6]]
输出：12
```


# 解题思路
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
<!-- Describe your approach to solving the problem. -->

1. Example
```
1  3  1
1  5  1
4  2  1
```

2. Calculate the first row & column minimum value
```
1  4  5
2  X  X
6  X  X
```

3. If we want to reach point A and got minimum value, We only need to check **min(4+A,2+A)** ; So A = 2 + 5 = 7

```
1  4  5
2  A  X
6  X  X
```

4. If we want to reach point B and got minimum value, We only need to check **min(5+B,6+B)** ; So B = 5 + 1 = 6

```
1  4  5
2  7  B
6  X  X
```

5. If we want to reach point C and got minimum value, We only need to check **min(7+C,6+C)** ; So C = 6 + 2 = 8

```
1  4  5
2  7  6
6  C  X
```

6. If we want to reach point D and got minimum value, We only need to check **min(6+D,8+D)** ; So D = 6 + 1 = 7

```
1  4  5
2  7  6
6  8  D
```

7. Return bottom right number; **grid[len(grid)-1,len(grid[0])-1]**
```
1  4  5
2  7  6
6  8  7
```

# Complexity
- Time complexity: O(m x n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: O(1)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
func minPathSum(grid [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	row := len(grid)
	column := len(grid[0])
	for c := 1; c < column; c++ {
		grid[0][c] += grid[0][c-1]
	}
	for r := 1; r < row; r++ {
		grid[r][0] += grid[r-1][0]
	}
	for r := 1; r < row; r++ {
		for c := 1; c < column; c++ {
			grid[r][c] += min(grid[r][c-1],grid[r-1][c])
		}
	}
	return grid[row-1][column-1]
}
```