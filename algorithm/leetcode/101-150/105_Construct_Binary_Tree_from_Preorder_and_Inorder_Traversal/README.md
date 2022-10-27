# 105. Construct Binary Tree from Preorder and Inorder Traversal

Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

### Example 1:

![image description](tree.jpeg)

```
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
```

### Example 2:

```
Input: preorder = [-1], inorder = [-1]
Output: [-1]
```


### Constraints:

Constraints:

* 1 <= preorder.length <= 3000
* inorder.length == preorder.length
* -3000 <= preorder[i], inorder[i] <= 3000
* preorder and inorder consist of unique values.
* Each value of inorder also appears in preorder.
* preorder is guaranteed to be the preorder traversal of the tree.
* inorder is guaranteed to be the inorder traversal of the tree.

### Translate:

> 105. 从前序与中序遍历序列构造二叉树
> 
> 给定两个整数数组preorder 和 inorder，其中preorder 是二叉树的先序遍历， inorder是同一棵树的中序遍历，请构造二叉树并返回其根节点。
> 
> 提示:
> 
> 1 <= preorder.length <= 3000
> 
> inorder.length == preorder.length
> 
> -3000 <= preorder[i], inorder[i] <= 3000
> 
> preorder和inorder均 无重复 元素
> 
> inorder均出现在preorder
> 
> preorder保证 为二叉树的前序遍历序列
> 
> inorder保证 为二叉树的中序遍历序列