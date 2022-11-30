# 289. Game of Life

According to Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

The board is made up of an m x n grid of cells, where each cell has an initial state: live (represented by a 1) or dead (represented by a 0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

* Any live cell with fewer than two live neighbors dies as if caused by under-population.
* Any live cell with two or three live neighbors lives on to the next generation.
* Any live cell with more than three live neighbors dies, as if by over-population.
* Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

The next state is created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur simultaneously. Given the current state of the m x n grid board, return the next state.


### Example 1:

![image description](grid1.jpeg)

```
Input: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
Output: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
```

### Example 2:

![image description](grid2.jpeg)

```
Input: board = [[1,1],[1,0]]
Output: [[1,1],[1,1]]
```

### Constraints:

* m == board.length
* n == board[i].length
* 1 <= m, n <= 25
* board[i][j] is 0 or 1.

### Follow up:

* Could you solve it in-place? Remember that the board needs to be updated simultaneously: You cannot update some cells first and then use their updated values to update other cells.
* In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems when the active area encroaches upon the border of the array (i.e., live cells reach the border). How would you address these problems?

### Translate:

> 289. 生命游戏

根据百度百科，生命游戏，简称为 生命 ，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。

给定一个包含 m × n个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态： 1 即为 活细胞 （live），或 0 即为 死细胞 （dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。给你 m x n 网格面板 board 的当前状态，返回下一个状态。

进阶：

* 你可以使用原地算法解决本题吗？请注意，面板上所有格子需要同时被更新：你不能先更新某些格子，然后使用它们的更新后的值再更新其他格子。
* 本题中，我们使用二维数组来表示面板。原则上，面板是无限的，但当活细胞侵占了面板边界时会造成问题。你将如何解决这些问题？


### 解题思路

|current|current number|neighbor count|calculate|How to get next state|How to get current state|
|:---|:---:|:---:|:---:|:---:|:---:|
| alive | 1 | <2     | N/A = 1   | 1 >> 1 = 0 | 1 & 1 = 1 |
| alive | 1 | 2 or 3 | 1 + 2 = 3 | 3 >> 1 = 1 | 3 & 1 = 1 |
| alive | 1 | > 3    | N/A = 1   | 1 >> 1 = 0 | 1 & 1 = 1 |
| dead  | 0 | 3      | 0 + 2 = 2 | 2 >> 1 = 1 | 2 & 1 = 0 |
| dead  | 0 | <2     | N/A = 0   | 0 >> 1 = 0 | 0 & 1 = 0 |

### 代码

```golang
func gameOfLife(board [][]int) {
	rowIndex, colIndex := len(board), len(board[0])
	for i := 0; i < rowIndex; i++ {
		for j := 0; j < colIndex; j++ {
			neighbor := neighborCount(board, i, j)
			if board[i][j] == 1 {
				if neighbor == 2 || neighbor == 3 {
					board[i][j] += 2
				}
			} else if board[i][j] == 0 {
				if neighbor == 3 {
					board[i][j] += 2
				}
			}
		}
	}
	for i := 0; i < rowIndex; i++ {
		for j := 0; j < colIndex; j++ {
			board[i][j] >>= 1
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func neighborCount(board [][]int, rowIndex, colIndex int) int {
	count := 0
	for i := max(0, rowIndex-1); i <= min(rowIndex+1, len(board)-1); i++ {
		for j := max(0, colIndex-1); j <= min(colIndex+1, len(board[0])-1); j++ {
			if i == rowIndex && j == colIndex {
				continue
			}
			if board[i][j] & 1 == 1 {
				count++
			}
		}
	}
	return count
}
```