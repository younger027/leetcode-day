package list

/**
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}

		fast = fast.Next
	}

	slow.Next = nil
	return head
}
