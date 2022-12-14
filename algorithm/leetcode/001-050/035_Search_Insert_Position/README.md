# 35. Search Insert Position

Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

### Example 1:
    
    Input: [1,3,5,6], 5
    Output: 2
    
### Example 2:
    
    Input: [1,3,5,6], 2
    Output: 1
    
### Example 3:
    
    Input: [1,3,5,6], 7
    Output: 4
    
### Example 4:
    
    Input: [1,3,5,6], 0
    Output: 0

# 35. 搜索插入位置

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

### 示例 1:
```
输入: nums = [1,3,5,6], target = 5
输出: 2
```

### 示例 2:
```
输入: nums = [1,3,5,6], target = 2
输出: 1
```

### 示例 3:

```
输入: nums = [1,3,5,6], target = 7
输出: 4
```

### 提示:

* 1 <= nums.length <= 10^4
* -10^4 <= nums[i] <= 10^4
* nums 为i无重复元素i的i升序i排列数组
* -10^4 <= target <= 10^4