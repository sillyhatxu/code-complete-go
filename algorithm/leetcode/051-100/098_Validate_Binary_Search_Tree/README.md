# 98. Validate Binary Search Tree

Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

* The left subtree of a node contains only nodes with keys less than the node's key.
* The right subtree of a node contains only nodes with keys greater than the node's key.
* Both the left and right subtrees must also be binary search trees.


### Example 1:

![image description](tree1.jpeg)

```
Input: root = [2,1,3]
Output: true
```

### Example 2:

![image description](tree2.jpeg)

```
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
```

### Constraints:

* The number of nodes in the tree is in the range [1, 104].
* -231 <= Node.val <= 231 - 1

### Translate:

> 98. 验证二叉搜索树
> 
> 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
> 
> 有效 二叉搜索树定义如下：
> 
> 节点的左子树只包含 小于 当前节点的数。
> 
> 节点的右子树只包含 大于 当前节点的数。
> 
> 所有左子树和右子树自身必须也是二叉搜索树。