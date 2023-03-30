# 152. Maximum Product Subarray

Given an integer array nums, find a subarray that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.

### Example 1:

```
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
```

### Example 2:

```
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
```

### Constraints:

* 1 <= nums.length <= 2 * 104
* -10 <= nums[i] <= 10
* The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

### Translate:

> 152. 乘积最大子数组

给你一个整数数组 nums，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个32-位 整数。

子数组 是数组的连续子序列。


示例 1:

输入: nums = [2,3,-2,4]
输出: 6
解释:子数组 [2,3] 有最大乘积 6。
示例 2:

输入: nums = [-2,0,-1]
输出: 0
解释:结果不能为 2, 因为 [-2,-1] 不是子数组。


提示:

1 <= nums.length <= 2 * 104
-10 <= nums[i] <= 10
nums 的任何前缀或后缀的乘积都 保证是一个 32-位 整数


### 解题思路

三种情况

设前边的乘积为x

1. x为正数，x * nums[i]是最大值
2. x为负数，x * nums[i]是最大值（-2 * -8 = 16负数的话，越小的值，乘积结果越大）
3. 当前nums[i]是最大值

比较3个值
currentMax = max(max*nums[i], min*nums[i], nums[i])
currentMin = min(max*nums[i], min*nums[i], nums[i])
max = currentMax
min = currentMin
res = max(res,currentMax)

### index

```
乘积；最大值；最大乘积
```