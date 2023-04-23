package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {

	//l1 := InitListNode([]int{1, 4, 5})
	//l2 := InitListNode([]int{1, 3, 4})
	//l3 := InitListNode([]int{2, 6})

	l1 := InitListNode([]int{2})
	l2 := InitListNode([]int{})
	l3 := InitListNode([]int{-1})

	arrayNode := []*ListNode{l1, l2, l3}
	ret := MergeKLists(arrayNode)
	for ret != nil {
		fmt.Println("node.Val:", ret.Val)
		ret = ret.Next
	}
}
