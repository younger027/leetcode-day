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
	InitHeap()
}
