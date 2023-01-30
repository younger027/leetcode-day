package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := InitListNode([]int{1, 2, 4})
	l2 := InitListNode([]int{1, 3, 4})

	//ret := MergeTwoLists(l1, l2)
	ret := NonRecursiveMergeTwoList(l1, l2)
	for ret != nil {
		fmt.Println("node.value:", ret.Val)
		ret = ret.Next
	}
}
