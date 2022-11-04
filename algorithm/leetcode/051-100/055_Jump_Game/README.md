# 55. Jump Game

You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.

### Example 1:

```
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

### Example 2:

```
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
```

### Constraints:

* 1 <= nums.length <= 104
* 0 <= nums[i] <= 105

### Translate:

> 55. 跳跃游戏
> 
> 给定一个非负整数数组nums ，你最初位于数组的 第一个下标 。
> 
> 数组中的每个元素代表你在该位置可以跳跃的最大长度。
> 
> 判断你是否能够到达最后一个下标。
> 
> 示例1：
> 
> 输入：nums = [2,3,1,1,4]
> 
> 输出：true
> 
> 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
> 
> 示例2：
> 
> 输入：nums = [3,2,1,0,4]
> 
> 输出：false
> 
> 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。

### 解题思路
举例一个长一点的数组

> nums = [2, 3, 1, 5, 4, 1, 2, 3, 2, 3, 1, 4, 9]

1. 起始位置 index = 0，有此可以得出，**最多**可以**往右跳2步**
```
[2, 3, 1, 5, 4, 1, 2, 3, 2, 3, 1, 4, 9]
 ↑
    ↑（如果只跳1步，则指向这里 nums[1] = 3）
       ↑ （如果跳2步，则指向这里 nums[2] = 1）
也就是说，当index = 0，我们最远到达距离增大了
[2, (3, 1), 5, 4, 1, 2, 3, 2, 3, 1, 4, 9]
 ↑
```

2. 之后再跳的话，能跳到哪里，取决于新增加的有边界内容，也就是上边的（3, 1）

2.1
```
[2, (3, 1), 5, 4, 1, 2, 3, 2, 3, 1, 4, 9]
     ↑ 如果我们选择从nums[1] = 3起跳，那么我们会将最大距离扩大到（1，5，4）
[2, 3, （1, 5, 4）, 1, 2, 3, 2, 3, 1, 4, 9]
    ↑
```

2.2
```
[2, (3, 1), 5, 4, 1, 2, 3, 2, 3, 1, 4, 9]
        ↑ 如果我们选择从nums[2] = 1起跳，那么我们会将最大距离扩大到（5）
[2, 3, 1, (5）, 4, 1, 2, 3, 2, 3, 1, 4, 9]
       ↑
```

3. 很显然2.1的有边界更大，我们取2.1的右边界，分别选择，(1, 5, 4)，向右扩大，右边界
```
[2, 3, (1, 5, 4), 1, 2, 3, 2, 3, 1, 4, 9]
```

4. 如果最后可以扩大的右边界（也就是最远到达距离），大于或等于nums的长度，则返回true，反之返回false

### 代码

```golang
func canJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		max = mathMax(max, i+nums[i])
	}
	return max >= - len(nums)-1
}

func mathMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```