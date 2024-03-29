# 378. Kth Smallest Element in a Sorted Matrix

Given an n x n matrix where each of the rows and columns is sorted in ascending order, return the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

You must find a solution with a memory complexity better than O(n2).

### Example 1:

```
Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
Output: 13
Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13
```

### Example 2:

```
Input: matrix = [[-5]], k = 1
Output: -5
```

### Constraints:

* n == matrix.length == matrix[i].length
* 1 <= n <= 300
* -109 <= matrix[i][j] <= 109
* All the rows and columns of matrix are guaranteed to be sorted in non-decreasing order.
* 1 <= k <= n2

### Follow up:

Could you solve the problem with a constant memory (i.e., O(1) memory complexity)?
Could you solve the problem in O(n) time complexity? The solution may be too advanced for an interview but you may find reading this paper fun.

### Translate:

> 378. 有序矩阵中第 K 小的元素

给你一个n x n矩阵matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。

你必须找到一个内存复杂度优于O(n2) 的解决方案。

示例 1：

输入：matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
输出：13
解释：矩阵中的元素为 [1,5,9,10,11,12,13,13,15]，第 8 小元素是 13

示例 2：

输入：matrix = [[-5]], k = 1
输出：-5

提示：

* n == matrix.length
* n == matrix[i].length
* 1 <= n <= 300
* -109 <= matrix[i][j] <= 109
* 题目数据 保证 matrix 中的所有行和列都按 非递减顺序 排列
* 1 <= k <= n2

进阶：

你能否用一个恒定的内存(即 O(1) 内存复杂度)来解决这个问题?
你能在 O(n) 的时间复杂度下解决这个问题吗?这个方法对于面试来说可能太超前了，但是你会发现阅读这篇文章（this paper）很有趣。