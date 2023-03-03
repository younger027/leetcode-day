package list

func IsRingList(node *ListNode) bool {
	if node == nil {
		return false
	}

	slow := node
	quick := node
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next

		if quick == slow {
			return true
		}
	}

	return false

}

func RingIntersectionNode(node *ListNode) *ListNode {
	if node == nil {
		return node
	}

	slow := node
	quick := node

	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next

		if quick == slow {
			return quick

		}
	}

	return nil
}

func RingLength(node *ListNode) int {
	if node == nil {
		return 0
	}

	slow := node
	quick := node

	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next

		if quick == slow {
			break
		}
	}

	count := 1
	quick = quick.Next
	for slow != quick {
		count += 1
		quick = quick.Next
	}

	return count
}
