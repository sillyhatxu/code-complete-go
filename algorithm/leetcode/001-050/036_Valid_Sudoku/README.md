# 36. Valid Sudoku

Determine if a `9 x 9` Sudoku board is valid. Only the filled cells need to be validated **according to the following rules**:

Each row must contain the digits `1-9` without repetition.
Each column must contain the digits `1-9` without repetition.
Each of the nine `3 x 3` sub-boxes of the grid must contain the digits 1-9 without repetition.

**Note:**

* A Sudoku board (partially filled) could be valid but is not necessarily solvable.
* Only the filled cells need to be validated according to the mentioned rules.


### Example 1:

![image description](Sudoku-by-L2G-20050714.svg.webp)

```
Input: board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
```

### Example 2:

```
Input: board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
```


### Constraints:

Constraints:

* board.length == 9
* board[i].length == 9
* board[i][j] is a digit 1-9 or '.'.

### Translate:

> 36. 有效的数独
> 请你判断一个9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字1-9在每一行只能出现一次。
数字1-9在每一列只能出现一次。
数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）


注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
空白格用'.'表示。

### Example 2:

```
输入：board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
输出：false
解释：除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。 但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。

```


### 解题思路

主要是找到grid index的计算方法

x(row % 3) + y(column % 3)最后得出如下表格 

| columns\rows   |      0      |  1 | 2 |
|----------|:-------------:|------:|------:|
| 0 |  0+0=0 | 0+1=1 | 0+2=2 |
| 1 |  1+0=1 | 1+1=2 | 1+2=3 |
| 2 |  2+0=2 | 2+1=3 | 2+2=4 |

实际需要的表格坐标

| columns\rows   |      0      |  1 | 2 |
|----------|:-------------:|------:|------:|
| 0 |  0 | 1 | 2 |
| 1 |  3 | 4 | 5 |
| 2 |  6 | 7 | 8 |
我们可以看出，第0行是符合要求的，所以此公式` x(row % 3) + y(column % 3) ` y = 1

再计算第二行

* x * 1+0 = 3
* x * 1+1 = 4
* x * 1+2 = 5

推出x = 3

所以表格坐标gridIndex = 3row%3 + column%3

### 代码

```golang
func isValidSudoku(board [][]byte) bool {
    rowMap := [9][9]bool{}
	columnMap := [9][9]bool{}
	gridMap := [9][9]bool{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			cell := string(board[i][j])
			if cell == "." {
				continue
			}
			val, _ := strconv.Atoi(cell)
			val--
			gridIndex := i/3*3 + j/3
			if rowMap[i][val] || columnMap[j][val] || gridMap[gridIndex][val] {
				return false
			}
			rowMap[i][val] = true
			columnMap[j][val] = true
			gridMap[gridIndex][val] = true
		}
	}
	return true
}
```