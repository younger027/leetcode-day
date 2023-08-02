package interview

import "fmt"

//算法面试题重刷
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func InitTreeNode(data []int, i int) *TreeNode {
	head := &TreeNode{
		Value: data[i],
		Left:  nil,
		Right: nil,
	}

	if i < len(data) && 2*i+1 < len(data) {
		head.Left = InitTreeNode(data, 2*i+1)
	}

	if i < len(data) && 2*i+2 < len(data) {
		head.Right = InitTreeNode(data, 2*i+2)
	}

	return head
}

//二叉树后序遍历-迭代
func BinartTreeBack(node *TreeNode) []int {
	var result []int
	if node == nil {
		return result
	}

	stack := make([]*TreeNode, 0)
	for node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		if node.Right != nil {
			node = node.Right
			continue
		}
		result = append(result, node.Value)
		node = stack[len(stack)-1]
	}

	return result
}

//递归法
func BinartTreeBack_2(node *TreeNode) {
	if node == nil {
		return
	}

	BinartTreeBack_2(node.Left)
	BinartTreeBack_2(node.Right)

	fmt.Println(node.Value)
}
