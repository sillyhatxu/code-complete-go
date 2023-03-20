# 236. Lowest Common Ancestor of a Binary Tree

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

### Example 1:

![image description](binarytree.png)

```
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.
```

### Example 2:

![image description](binarytree.png)

```
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.
```

### Example 3:

```
Input: root = [1,2], p = 1, q = 2
Output: 1
```

### Constraints:

* The number of nodes in the tree is in the range [2, 105].
* -109 <= Node.val <= 109
* All Node.val are unique.
* p != q
* p and q will exist in the tree.

### Translate:

> 236. 二叉树的最近公共祖先

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

`百度百科`中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”


### 示例 1：

```
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
```

### 示例 2：

```
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
```

### 示例 3：

```
输入：root = [1,2], p = 1, q = 2
输出：1
```

### 提示：

* 树中节点数目在范围 [2, 105] 内。
* -109 <= Node.val <= 109
* 所有 Node.val 互不相同 。
* p != q
* p 和 q 均存在于给定的二叉树中。


### 解题思路
Recurtion; Time: O(n) ; Space: O(1)

一共有三种情况

![image description](binarytree.png)

1. p,q 位于当前root节点的两侧，所以返回当前root节点，如
```
p=5;q=1 return 3;
p=0;q=5 return 1;
```   

2. p,q 位于当前root节点同侧，但是p在上，q在下，所以返回p，如
```
p=5;q=4 return 5;
```  

3. p,q 位于当前root节点同侧，但是q在上，p在下，所以返回q，如
```
p=7;q=2 return 2;
```  

### 代码

```golang
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
```

### index

```
找到最近公共祖先，公共root
```