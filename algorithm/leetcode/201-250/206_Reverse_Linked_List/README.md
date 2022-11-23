# 206. Reverse Linked List

Given the head of a singly linked list, reverse the list, and return the reversed list.

### Example 1:

![image description](rev1ex1.jpeg)

```
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
```

### Example 2:

![image description](rev1ex2.jpeg)

```
Input: head = [1,2]
Output: [2,1]
```

### Example 3:

```
Input: head = []
Output: []
```

### Constraints:

* The number of nodes in the list is the range [0, 5000].
* -5000 <= Node.val <= 5000

### Translate:

> 206. 反转链表

给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？





### 解题思路

Minimum Memory; iterative and recursive

### 代码

#### iterative
```golang
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
```

#### recursive

```golang
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	res := reverseList(head.Next)
	nextHead := head.Next
	nextHead.Next = head
	head.Next = nil
	return res
}
```