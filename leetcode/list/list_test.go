package list

import (
	"fmt"
	"testing"
)

func InitListNode(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  data[0],
		Next: nil,
	}

	start := head
	for i := 1; i < len(data); i++ {
		item := ListNode{
			Val:  data[i],
			Next: nil,
		}

		head.Next = &item
		head = head.Next

	}

	return start
}

func TestPartition(t *testing.T) {
	head := InitListNode([]int{1, 4, 3, 2, 5, 2})
	ret := Partition(head, 3)
	for ret != nil {
		fmt.Println("node.value:", ret.Val)
		ret = ret.Next
	}
}

func TestInitHeap(t *testing.T) {
	//InitHeap()

	l1 := InitListNode([]int{2, 3, 5})
	l2 := InitListNode([]int{1, 6})
	l3 := InitListNode([]int{4, 7})

	ret := MergeKLists([]*ListNode{l1, l2, l3})
	for ret != nil {
		fmt.Println("node.value:", ret.Val)
		ret = ret.Next
	}
}

func TestMiddleNode(t *testing.T) {
	l1 := InitListNode([]int{1, 2, 3, 4, 5, 6})

	t.Log(MiddleNode(l1).Val)
}

func TestIsRingList(t *testing.T) {
	l1 := InitListNode([]int{1, 2, 3, 4, 5, 6})

	head := l1
	l1.Next.Next.Next.Next.Next.Next = head.Next.Next.Next

	t.Log(IsRingList(l1))

	t.Log(RingIntersectionNode(l1).Val)

	t.Log(RingLength(l1))
}

func TestGetIntersectionNode(t *testing.T) {
	l1 := InitListNode([]int{1, 2, 3, 4, 5, 6, 10})
	//l2 := InitListNode([]int{7, 8, 9})

	//l2.Next.Next.Next = l1.Next.Next.Next

	//t.Log(GetIntersectionNode(l1, l2))

	node := ReverseList(l1)
	for node != nil {
		fmt.Println("node---", node.Val)
		node = node.Next
	}
}
