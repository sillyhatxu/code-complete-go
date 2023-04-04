# 416. Partition Equal Subset Sum

Given a non-empty array nums containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

### Example 1:

```
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].
```

### Example 2:

```
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.
```

### Constraints:

* 1 <= nums.length <= 200
* 1 <= nums[i] <= 100

### Translate:

> 416. 分割等和子集

给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

### 示例 1：

```
输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。
```

### 示例 2：

```
输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。
```

### 提示：

* 1 <= nums.length <= 200
* 1 <= nums[i] <= 100


### 解题思路

```
1. sum必须是偶数才能拆分
2. dp[i] 表示能否从数组中选取若干个元素，其元素和等于i
```

### index

```
数组；和；拆分
```