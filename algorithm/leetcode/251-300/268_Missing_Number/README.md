# 268. Missing Number

Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

### Example 1:

```
Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
```

### Example 2:

```
Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.
```

### Example 3:

```
Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.
```

### Constraints:

* n == nums.length
* 1 <= n <= 104
* 0 <= nums[i] <= n
* All the numbers of nums are unique.

### Follow up: 

Could you implement a solution using only O(1) extra space complexity and O(n) runtime complexity?


### Translate:

> 268. 丢失的数字

给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

示例 1：

输入：nums = [3,0,1]
输出：2
解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。
示例 2：

输入：nums = [0,1]
输出：2
解释：n = 2，因为有 2 个数字，所以所有的数字都在范围 [0,2] 内。2 是丢失的数字，因为它没有出现在 nums 中。
示例 3：

输入：nums = [9,6,4,2,3,5,7,0,1]
输出：8
解释：n = 9，因为有 9 个数字，所以所有的数字都在范围 [0,9] 内。8 是丢失的数字，因为它没有出现在 nums 中。
示例 4：

输入：nums = [0]
输出：1
解释：n = 1，因为有 1 个数字，所以所有的数字都在范围 [0,1] 内。1 是丢失的数字，因为它没有出现在 nums 中。

提示：

* n == nums.length
* 1 <= n <= 104
* 0 <= nums[i] <= n
* nums 中的所有数字都 独一无二

进阶：你能否实现线性时间复杂度、仅使用额外常数空间的算法解决此问题?



### 解题思路
Sample solution; Time: O(n); Space: O(1)

result = len(nums)

||||||
|:---:|:---:|:---:|:---:|---:|
| input | 0 | 1 | 2 |  |
| i | 0 | 1 | 2 |  |
| result | 0^0^3=3 | 1^1^3=3 | 2^2^3=3 | 3 |


||||||||
|:---:|:---:|:---:|:---:|:---:|:---:|---:|
| input | 0 | 1 | 2 | 3 | 5 |   |
| i     | 0 | 1 | 2 | 3 | 4 |   |
| result | 0^0^5=5 | 1^1^5=5 | 2^2^5=5 | 3^3^5=5 | 5^4^5=4 |  4  |

### 代码

```golang
func missingNumber(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res ^= nums[i] ^ i
	}
	return res
}
```