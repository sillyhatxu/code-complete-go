# 169. Majority Element

Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

### Example 1:

```
Input: nums = [3,2,3]
Output: 3
```

### Example 2:

```
Input: nums = [2,2,1,1,1,2,2]
Output: 2
```

### Constraints:

* n == nums.length
* 1 <= n <= 5 * 104
* -109 <= nums[i] <= 109

### Translate:

> 169. 多数元素

给定一个大小为 n 的数组nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。


### 解题思路

因为多数元素出现次数一定大于⌊ n/2 ⌋

所以对不同元素计算出现次数即可
```
如：12344564474849
```
|index|nums|num|count|
|---|:---:|:---:|:---:|
| 0 | 1 | 1 | 1 |
| 1 | 2 | 1 | 0 |
| 2 | 3 | 3 | 1 |
| 3 | 4 | 3 | 0 |
| 4 | 4 | 4 | 1 |
| 5 | 5 | 4 | 0 |
| 6 | 6 | 6 | 1 |
| 7 | 4 | 6 | 0 |
| 8 | 4 | 4 | 1 |
| 9 | 7 | 4 | 0 |
| 10 | 4 | 4 | 1 |
| 11 | 8 | 4 | 0 |
| 12 | 4 | 4 | 1 |
| 13 | 9 | 4 | 0 |





