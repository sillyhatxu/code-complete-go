# 73. Set Matrix Zeroes

Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.


### Example 1:

![image description](mat1.jpeg)

```
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]
```

### Example 2:

![image description](mat2.jpeg)

```
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
```

### Constraints:

* m == matrix.length
* n == matrix[0].length
* 1 <= m, n <= 200
* -2^31 <= matrix[i][j] <= 2^31 - 1

### Follow up:

* A straightforward solution using O(mn) space is probably a bad idea.
* A simple improvement uses O(m + n) space, but still not the best solution.
* Could you devise a constant space solution?

### Translate:

> 73. 矩阵置零
> 
> 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

进阶：

一个直观的解决方案是使用 O(mn)的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m+n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个仅使用常量空间的解决方案吗？


### 解题思路

假设输入一个大的二维数组

||||||
|---|:---:|:---:|:---:|---:|
| 1 | 1 | 1 | 1 | 0 |
| 1 | 2 | 2 | 2 | 2 |
| 1 | 2 | 2 | 2 | 2 |
| 1 | 2 | 0 | 2 | 2 |
| 1 | 2 | 2 | 0 | 2 |

1. 遍历第0行（从第1元素开始），查看是否有0如果存在，zeroInRow = true
2. 遍历第0列（从第1元素开始），查看是否有0如果存在，zeroInColumn = true
3. 遍历除掉第0行和第0例的内容，也就是此例中，2的区域，查看是否有0，如果有零，则设置二维数组的0行或0列为0
4. 由123可知，直接遍历整个二维数组，操作后得出如下
* zeroInRow = true
* zeroInColumn = false

||||||
|:---:|:---:|:---:|:---:|:---:|
| 0 | 1 | 0 | 0 | 0 |
| 1 | 2 | 2 | 2 | 2 |
| 1 | 2 | 2 | 2 | 2 |
| 0 | 2 | 0 | 2 | 2 |
| 0 | 2 | 2 | 0 | 2 |

5. 遍历第0行，判断是否为0，如果为0，则设置当前列其他元素为0
6. 遍历第0列，判断是否为0，如果为0，则设置当前行其他元素为0
* zeroInRow = true
* zeroInColumn = false

||||||
|:---:|:---:|:---:|:---:|:---:|
| 0 | 1 | 0 | 0 | 0 |
| 1 | 2 | 0 | 0 | 0 |
| 1 | 2 | 0 | 0 | 0 |
| 0 | 0 | 0 | 0 | 0 |
| 0 | 0 | 0 | 0 | 0 |

7. 判断zeroInRow，如果为true，则把第0行，全部改为0
8. 判断zeroInColumn，如果为true，则把第0列，全部改为0

||||||
|:---:|:---:|:---:|:---:|:---:|
| 0 | 0 | 0 | 0 | 0 |
| 1 | 2 | 0 | 0 | 0 |
| 1 | 2 | 0 | 0 | 0 |
| 0 | 0 | 0 | 0 | 0 |
| 0 | 0 | 0 | 0 | 0 |

### 代码

```golang
func setZeroes(matrix [][]int) {
	zeroInRow, zeroInColumn := false, false
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				if row == 0 {
					zeroInRow = true
				}
				if col == 0 {
					zeroInColumn = true
				}
				matrix[0][col], matrix[row][0] = 0, 0
			}
		}
	}
	for col := 1; col < len(matrix[0]); col++ {
		if matrix[0][col] == 0 {
			for i := 1; i < len(matrix); i++ {
				matrix[i][col] = 0
			}
		}
	}
	for row := 1; row < len(matrix); row++ {
		if matrix[row][0] == 0 {
			for col := 1; col < len(matrix[0]); col++ {
				matrix[row][col] = 0
			}
		}
	}
	if zeroInRow {
		for col := 0; col < len(matrix[0]); col++ {
			matrix[0][col] = 0
		}
	}
	if zeroInColumn {
		for row := 0; row < len(matrix); row++ {
			matrix[row][0] = 0
		}
	}
}
```