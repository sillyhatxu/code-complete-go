# 136. Single Number

Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

### Example 1:

```
Input: nums = [2,2,1]
Output: 1
```

### Example 2:

```
Input: nums = [4,1,2,1,2]
Output: 4
```

### Example 3:

```
Input: nums = [1]
Output: 1
```

### Constraints:

* 1 <= nums.length <= 3 * 104
* -3 * 104 <= nums[i] <= 3 * 104
* Each element in the array appears twice except for one element which appears only once.

### Translate:

> 136. 只出现一次的数字

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？


### 解题思路
首先，此题中传入参数是成对出现的，只有一个元素不是成对出现。所以可以使用 **异或(XOR)** 方式计算

#### 复习 **异或(XOR)**


#### 3 ^ 4 = 7
```
	  1 1
	1 0 0

=   1 1 1
= 2^2 + 2^1 + 2^0 = 4 + 2 + 1 = 7
```

#### 5 ^ 5 = 0
```
	1 0 1
	1 0 1
    
=   0 0 0
=   0
```
#### 传入参数 [2, 1, 2]
```
	  1 0
^       1
=     1 1
^     1 0    
=     0 1
= 1
```

### 代码

```golang
func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
```