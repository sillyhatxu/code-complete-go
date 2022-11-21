# 189. Rotate Array

Given an array, rotate the array to the right by k steps, where k is non-negative.

### Example 1:

```
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
```

### Example 2:

```
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation: 
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]

```

### Constraints:

* 1 <= nums.length <= 105
* -231 <= nums[i] <= 231 - 1
* 0 <= k <= 105

### Follow up:

* Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
* Could you do it in-place with O(1) extra space?

### Translate:

> 189. 轮转数组

给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:

```
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
```

示例2:

```
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]
```

### 提示：

* 1 <= nums.length <= 105
* -231 <= nums[i] <= 231 - 1
* 0 <= k <= 105


### 进阶：

* 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
* 你可以使用空间复杂度为O(1) 的原地算法解决这个问题吗？


### 解题思路

翻转 1，2，3，4，5，6，7        k=3
```
1. 全翻转
7，6，5，4，3，2，1

2.  以k为中线，分别翻转左右
7   6   5   |   4   3   2   1
5   6   7       1   2   3   7

得到答案
```


### 代码

```golang
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
```