package list

import (
	"container/heap"
	"fmt"
)

/*
23. 合并K个升序链表
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]


提示：
k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4
*/

//// An IntHeap is a min-heap of ints.
//type IntHeap []int
//
//func (h IntHeap) Len() int           { return len(h) }
//func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
//func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
//func (h *IntHeap) Push(x any) {
//	// Push and Pop use pointer receivers because they modify the slice's length,
//	// not just its contents.
//	*h = append(*h, x.(int))
//}
//
//func (h *IntHeap) Pop() any {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}

type NodeHeap struct {
	data []*ListNode
}

func (h NodeHeap) Len() int           { return len(h.data) }
func (h NodeHeap) Less(i, j int) bool { return h.data[i].Val < h.data[j].Val }
func (h NodeHeap) Swap(i, j int)      { h.data[i].Val, h.data[j].Val = h.data[j].Val, h.data[i].Val }

func (h *NodeHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.data = append(h.data, x.(*ListNode))
}

func (h *NodeHeap) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func InitHeap() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  4,
		Next: nil,
	}

	l3 := &ListNode{
		Val:  5,
		Next: nil,
	}

	l4 := &ListNode{
		Val:  3,
		Next: nil,
	}

	h := &NodeHeap{
		data: []*ListNode{l1, l2, l3},
	}

	heap.Init(h)
	heap.Push(h, l4)

	fmt.Printf("minimum: %d\n", h.data[0].Val)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}

//solution 1 base on min heap

//func mergeKLists(lists []*ListNode) *ListNode {
//	h := &ListNode{}
//	heap.Init(h)
//
//	for _, list := range lists {
//		if list != nil {
//			heap.Push(h, list.Val)
//		}
//	}
//
//	result := &ListNode{
//		Val:  0,
//		Next: nil,
//	}
//
//}
