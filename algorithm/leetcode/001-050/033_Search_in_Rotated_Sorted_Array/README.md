# 33. Search in Rotated Sorted Array

There is an integer array `nums` sorted in ascending order (with **distinct** values).

Prior to being passed to your function, ``nums`` is **possibly rotated** at an unknown pivot index `k (1 <= k < nums.length)` such that the resulting array is `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` **(0-indexed)**. For example, `[0,1,2,4,5,6,7]` might be rotated at pivot index `3` and become `[4,5,6,7,0,1,2]`.

Given the array `nums` after the possible rotation and an integer `target`, return the index of `target` if it is in `nums`, or `-1` if it is not in `nums`.

You must write an algorithm with `O(log n)` runtime complexity.


### Example 1:

```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

### Example 2:

```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

### Example 3:

```
Input: nums = [1], target = 0
Output: -1
```

### Constraints:

* 1 <= nums.length <= 5000
* -104 <= nums[i] <= 104
* All values of nums are unique.
* nums is an ascending array that is possibly rotated.
* -104 <= target <= 104

### Translate:

> 33. 搜索旋转排序数组
> 
> 整数数组 nums 按升序排列，数组中的值 互不相同 。
> 
> 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
> 
> 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
> 
> 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
> 
> 提示：
> 
> 1 <= nums.length <= 5000
> 
> -104 <= nums[i] <= 104
> 
> nums 中的每个值都 独一无二
> 
> 题目数据保证 nums 在预先未知的某个下标上进行了旋转
> 
> -104 <= target <= 104



### 解题思路

#### Time: O(logn) Space: O(1) 0ms 2 binary searches

分两个步骤

```
数组[7,8,9,0,1,2,3,4,5,6]
   [7,8,9,  0,1,2,3,4,5,6]
找到这个位置↑
```

1. 二分查找，找到破坏递增规则的index
2. 判断target在左边区域还是右边区域，然后去对应区域继续二分查找


### 代码

```golang
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 && nums[0] == target {
		return 0
	} else if len(nums) == 1 && nums[0] != target {
		return -1
	}
	pivot := findPivotIndex(nums)
	if pivot >= 0 && target >= nums[0] && target <= nums[pivot] {
		return binarySearch(nums, 0, pivot, target)
	} else {
		return binarySearch(nums, pivot+1, len(nums)-1, target)
	}
}

func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func findPivotIndex(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return len(nums) - 1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			return mid
		} else if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}

```