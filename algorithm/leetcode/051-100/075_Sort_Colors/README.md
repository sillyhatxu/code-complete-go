# 75. Sort Colors

Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.


### Example 1:

```
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

### Example 2:

```
Input: nums = [2,0,1]
Output: [0,1,2]
```

### Constraints:

* n == nums.length
* 1 <= n <= 300
* nums[i] is either 0, 1, or 2.

### Follow up: 

Could you come up with a one-pass algorithm using only constant extra space?

### Translate:

> 75. 颜色分类
> 
> 给定一个包含红色、白色和蓝色、共n 个元素的数组nums，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
> 
> 我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。
> 
> 必须在不使用库的sort函数的情况下解决这个问题。
> 
> 进阶：
> 
> 你可以不使用代码库中的排序函数来解决这道题吗？
> 
> 你能想出一个仅使用常数空间的一趟扫描算法吗？

### 解题思路
快排的变种，因为只有0,1,2三个数字，所以不需要遍历分区。

### 代码

```golang
func sortColors(nums []int)  {
	i, p0, p2 := 0, 0, len(nums)-1
	for i <= p2 {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++
			continue
		} else if nums[i] == 1 {
			i++
			continue
		} else if nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
			continue
		}
	}
}
```
