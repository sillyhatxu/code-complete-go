# 34. Find First and Last Position of Element in Sorted Array

Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target
value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

### Example 1:

```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```

### Example 2:

```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```

### Example 3:

```
Input: nums = [], target = 0
Output: [-1,-1]
```

### Constraints:

* 0 <= nums.length <= 105
* -109 <= nums[i] <= 109
* nums is a non-decreasing array.
* -109 <= target <= 109

### Translate:

> 34. 在排序数组中查找元素的第一个和最后一个位置
>
> 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
>
> 如果数组中不存在目标值 target，返回[-1, -1]。
>
> 你必须设计并实现时间复杂度为O(log n)的算法解决此问题。
>
> 0 <= nums.length <= 105
>
> -109<= nums[i]<= 109
>
> nums是一个非递减数组
>
> -109<= target<= 109