# 112. Path Sum

Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.

A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

### Example 1:

![image description](108.1.1.jpeg)

```
Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:
```
![image description](108.1.2.jpeg)

### Example 2:

![image description](108.2.jpeg)

```
Input: root = [1,2,3], targetSum = 5
Output: false
Explanation: There two root-to-leaf paths in the tree:
(1 --> 2): The sum is 3.
(1 --> 3): The sum is 4.
There is no root-to-leaf path with sum = 5.
```


### Constraints:

* 1 <= nums.length <= 104
* -104 <= nums[i] <= 104
* nums is sorted in a strictly increasing order.

### Translate:

> 108. 将有序数组转换为二叉搜索树

给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。


### Recursion 0 ms 100%; Time: O(n); Space: O(logn)

```golang
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(&nums, 0, len(nums)-1)
}

func helper(nums *[]int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	return &TreeNode{
		Val:   (*nums)[mid],
		Left:  helper(nums, left, mid-1),
		Right: helper(nums, mid+1, right),
	}
}
```