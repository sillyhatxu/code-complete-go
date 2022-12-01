# 324. Wiggle Sort II

Given an integer array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....

You may assume the input array always has a valid answer.

### Example 1:

```
Input: nums = [1,5,1,1,6,4]
Output: [1,6,1,5,1,4]
Explanation: [1,4,1,5,1,6] is also accepted.
```

### Example 2:

```
Input: nums = [1,3,2,2,3,1]
Output: [2,3,1,3,1,2]
```

### Constraints:

* 1 <= nums.length <= 5 * 104
* 0 <= nums[i] <= 5000
* It is guaranteed that there will be an answer for the given input nums.

### Follow Up

Can you do it in O(n) time and/or in-place with O(1) extra space?

### Translate:

> 324. 摆动排序 II

给你一个整数数组nums，将它重新排列成nums[0] < nums[1] > nums[2] < nums[3]...的顺序。

你可以假设所有输入数组都可以得到满足题目要求的结果。

示例 1：

输入：nums = [1,5,1,1,6,4]
输出：[1,6,1,5,1,4]
解释：[1,4,1,5,1,6] 同样是符合题目要求的结果，可以被判题程序接受。

示例 2：

输入：nums = [1,3,2,2,3,1]
输出：[2,3,1,3,1,2]

提示：

1 <= nums.length <= 5 * 104
0 <= nums[i] <= 5000
题目数据保证，对于给定的输入 nums ，总能产生满足题目要求的结果

进阶：你能用O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？


### 解题思路


### Time: O(n); Space: O(n)

```golang
func wiggleSort(nums []int)  {
    sort.Ints(nums)
	temp := make([]int, len(nums), len(nums))
	for i := range nums {
		temp[i] = nums[i]
	}
	midIndex := (len(nums)+1)/2 - 1
	endIndex := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = temp[midIndex]
			midIndex--
		} else {
			nums[i] = temp[endIndex]
			endIndex--
		}
	}
}
```