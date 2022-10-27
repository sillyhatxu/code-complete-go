# 106. Construct Binary Tree from Inorder and Postorder Traversal

Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.
### Example 1:

![image description](tree.jpeg)

```
Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
Output: [3,9,20,null,null,15,7]
```

### Example 2:

```
Input: inorder = [-1], postorder = [-1]
Output: [-1]
```


### Constraints:

Constraints:

* 1 <= inorder.length <= 3000
* postorder.length == inorder.length
* -3000 <= inorder[i], postorder[i] <= 3000
* inorder and postorder consist of unique values.
* Each value of postorder also appears in inorder.
* inorder is guaranteed to be the inorder traversal of the tree.
* postorder is guaranteed to be the postorder traversal of the tree.

### Translate:

> 106. 从中序与后序遍历序列构造二叉树
> 
> 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗二叉树。
> 
> 提示:
>
>  1 <= inorder.length <= 3000
> 
>  postorder.length == inorder.length
> 
>  -3000 <= inorder[i], postorder[i] <= 3000
> 
>  inorder和postorder都由 不同 的值组成
> 
>  postorder中每一个值都在inorder中
> 
>  inorder保证是树的中序遍历
> 
>  postorder保证是树的后序遍历