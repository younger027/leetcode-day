package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
复杂度分析

如何计算递归的时间复杂度和空间复杂度呢？ 力扣对此进行了 详细介绍 ，其中时间复杂度可以这样计算：

给出一个递归算法，其时间复杂度O(T) 通常是递归调用的数量（记作R）和计算的时间复杂度的乘积（表示为 O(s)的乘积：
O(T)=R∗O(s)

时间复杂度：O(m+n)
m,n为l1, l2的元素个数。递归函数每次去掉一个元素，直到两个链表都为空，因此需要调用 R=O(m+n) 次。而在递归函数中我们只进行了 next 指针的赋值操作，复杂度为 O(1)，
故递归的总时间复杂度为 O(T)=R∗O(1)=O(m+n) 。


空间复杂度：O(m+n)
对于递归调用 self.mergeTwoLists()，当它遇到终止条件准备回溯时，已经递归调用了m+n次，使用了 m+n 个栈帧，
故最后的空间复杂度为 O(m+n){\mathcal{O}}(m + n)O(m+n)。

*/
/*
recursive solution 递归解法，需要寻找子问题和结束条件
子问题：合并剩下的节点
结束的条件：两个链表走到尾部
*/

//1 2 4, 1 3 4
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists(list1, list2.Next)
		return list2
	}
}

/*
时间复杂度分析
时间复杂度：O(m+n),每次循环只会将一个节点放到链表中，总共m+n个节点
空间复杂度：只创建了一个空的头结点，O(1)
*/
func NonRecursiveMergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	head := new(ListNode)
	prev := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}

		prev = prev.Next
	}

	if list1 == nil {
		prev.Next = list2
	}

	if list2 == nil {
		prev.Next = list1
	}

	return head.Next
}
