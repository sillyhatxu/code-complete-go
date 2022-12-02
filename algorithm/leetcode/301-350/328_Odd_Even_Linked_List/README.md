# 328. Odd Even Linked List

Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

The first node is considered odd, and the second node is even, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in O(1) extra space complexity and O(n) time complexity.

### Example 1:

![image description](oddeven-linked-list.jpeg)

```
Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]
```

### Example 2:

![image description](oddeven2-linked-list.jpeg)

```
Input: head = [2,1,3,5,6,4,7]
Output: [2,3,6,7,1,5,4]
```

### Constraints:

* The number of nodes in the linked list is in the range [0, 104].
* -106 <= Node.val <= 106

### Translate:

> 328. 奇偶链表

给定单链表的头节点head，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。

第一个节点的索引被认为是 奇数，第二个节点的索引为偶数，以此类推。

请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。

你必须在O(1)的额外空间复杂度和O(n)的时间复杂度下解决这个问题。