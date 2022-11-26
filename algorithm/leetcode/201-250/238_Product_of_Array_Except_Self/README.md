# 238. Product of Array Except Self

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.


### Example 1:

```
Input: nums = [1,2,3,4]
Output: [24,12,8,6]
```

### Example 2:

```
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
```

### Constraints:

* 2 <= nums.length <= 105
* -30 <= nums[i] <= 30
* The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

### Follow up

Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)

### Translate:

> 238. 除自身以外数组的乘积

给你一个整数数组nums，返回 数组answer，其中answer[i]等于nums中除nums[i]之外其余各元素的乘积。

题目数据 保证 数组nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。

请不要使用除法，且在O(n) 时间复杂度内完成此题。



示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]


提示：

2 <= nums.length <= 105
-30 <= nums[i] <= 30
保证 数组nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内

进阶：你可以在 O(1)的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）


### 解题思路

此题目为：前缀和/积系列,与之相对的套路就是差分，进阶就是二维前缀和/积，差分

```
传入数组[1,2,3,4]
输出数组[24,12,8,6]

24 =     2 * 3 * 4
12 = 1 *     3 * 4
 8 = 1 * 2 *     4
 6 = 1 * 2 * 3    
 
我们可以从前往后依次乘积，再从后往前依次乘积，得到两个数组
left = [     ][ 1 ][1*2][1*2*3]
right= [2*3*4][3*4][ 4 ][     ]
对比上边每位的乘积，我们就可以得出，上下两个数组乘积，得出来的新数组，就是答案
但是 left的第一位和right的最后一位分别为空，不好计算，所以我们使用1作为占位符（因为1乘任何数字都是1，不会改变答案）

left = [    1    ] [ (1)*1 ] [(1)*1*2] [(1)*1*2*3]
right= [(1)*2*3*4] [(1)*3*4] [ (1)*4 ] [    1    ]

answer=[    24   ] [   12  ] [   8   ] [    6    ]
```
### 代码

```golang
func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length, length)
	res[0] = 1
	current := nums[0]
	for i := 1; i < length; i++ {
		res[i] = current
		current *= nums[i]
	}
	current = nums[length-1]
	for i := length - 2; i >= 0; i-- {
		res[i] *= current
		current *= nums[i]
	}
	return res
}
```