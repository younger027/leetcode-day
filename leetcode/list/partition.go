package list

/*
86. 分隔链表
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。

示例 1：
输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]

示例 2：
输入：head = [2,1], x = 2
输出：[1,2]

提示：
链表中节点的数目在范围 [0, 200] 内
-100 <= Node.val <= 100
-200 <= x <= 200
*/

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

func Partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	p1 := &ListNode{
		Val:  0,
		Next: nil,
	}

	p2 := &ListNode{
		Val:  0,
		Next: nil,
	}

	result := p1
	result2 := p2
	p := head
	for p != nil {
		if p.Val >= x {
			p2.Next = p
			p2 = p2.Next
		} else {
			p1.Next = p
			p1 = p1.Next
		}

		temp := p.Next
		p.Next = nil
		p = temp
	}

	p1.Next = result2.Next

	return result.Next
}
