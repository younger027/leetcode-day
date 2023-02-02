package leetcode

//func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	if list1 == nil {
//		return list2
//	}
//
//	if list2 == nil {
//		return list1
//	}
//
//	if list1.Val <= list2.Val {
//		list1.Next = MergeTwoLists(list1.Next, list2)
//		return list1
//	} else {
//		list2.Next = MergeTwoLists(list1, list2.Next)
//		return list2
//	}
//}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	var prev *ListNode
	length := len(lists)
	if length == 0 {
		return head
	}

	if length == 1 {
		return lists[0]
	}

	left := 0
	for left < length {
		for left < length-1 && lists[left] == nil {
			left++
		}

		if left < length {
			prev = NonRecursiveMergeTwoList(prev, lists[left])
			head = prev
		}

		left += 1
	}

	return head
}
