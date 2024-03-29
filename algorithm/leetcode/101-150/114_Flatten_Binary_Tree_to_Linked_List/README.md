# 114. Flatten Binary Tree to Linked List

Given the root of a binary tree, flatten the tree into a "linked list":

* The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
* The "linked list" should be in the same order as a pre-order traversal of the binary tree.

### Example 1:

![image description](flaten.jpeg)

```
Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]
```

### Example 2:

```
Input: root = []
Output: []
```

### Example 3:

```
Input: root = [0]
Output: [0]
```

### Constraints:

* The number of nodes in the tree is in the range [0, 2000].
* -100 <= Node.val <= 100

### Follow up: 

Can you flatten the tree in-place (with O(1) extra space)?


### Translate:

> 114. 二叉树展开为链表

给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

### 示例 1：

![image description](flaten.jpeg)

```
输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]
```

### 示例 2：

```
输入：root = []
输出：[]
```

### 示例 3：

```
输入：root = [0]
输出：[0]
```

### 提示：

* 树中结点数在范围 [0, 2000] 内
* -100 <= Node.val <= 100


# 解题思路
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity: O(n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: O(1)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var stack []*TreeNode
	stack = append([]*TreeNode{root}, stack...)
	for len(stack) != 0 {
		curr := stack[0]
		stack = stack[1:]
		if curr.Right != nil {
			stack = append([]*TreeNode{curr.Right}, stack...)
		}
		if curr.Left != nil {
			stack = append([]*TreeNode{curr.Left}, stack...)
		}
		if len(stack) != 0 {
			curr.Right = stack[0]
		}
		curr.Left = nil
	}
}
```