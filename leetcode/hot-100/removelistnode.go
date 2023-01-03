package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	newHead := &ListNode{
		Val:  0,
		Next: head,
	}

	cur, quick := newHead, head

	i := 0
	for i < n {
		quick = quick.Next
		i++

	}

	for quick != nil {
		cur = cur.Next
		quick = quick.Next
	}

	cur.Next = cur.Next.Next

	return newHead.Next
}
