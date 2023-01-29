package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//1 2 4, 1 3 4
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	ret := list1
	l1 := list1
	p2, n2 := list2, list2
	for l1 != nil || n2 != nil {
		if l1 == nil {
			l1.Next = n2
		}

		if n2 == nil {
			break
		}

		if n2.Val <= l1.Val {
			p2.Next = n2.Next
			n2.Next = l1
			n2 = p2.Next

			l1 = l1.Next
		} else {
			p2 = n2
			n2 = n2.Next
		}
	}

	return ret
}
