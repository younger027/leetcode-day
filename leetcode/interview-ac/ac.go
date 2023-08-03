package interview

import (
	"container/list"
	"fmt"
)

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

	if 2*i+1 < len(data) {
		head.Left = InitTreeNode(data, 2*i+1)
	}

	if 2*i+2 < len(data) {
		head.Right = InitTreeNode(data, 2*i+2)
	}

	return head
}

func BinartTreeLevel(node *TreeNode) []int {
	var result []int
	if node == nil {
		return result
	}

	queue := new(list.List)
	queue.PushBack(node)
	for queue.Len() > 0 {
		tmpNode := queue.Remove(queue.Front()).(*TreeNode)
		result = append(result, tmpNode.Value)

		if tmpNode.Left != nil {
			queue.PushBack(tmpNode.Left)
		}

		if tmpNode.Right != nil {
			queue.PushBack(tmpNode.Right)
		}
	}

	return result
}

//二叉树先序遍历-迭代-中左右
func BinartTreePre(node *TreeNode) []int {
	var result []int
	if node == nil {
		return result
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, node)
	for len(stack) > 0 {
		topNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, topNode.Value)

		if topNode.Right != nil {
			stack = append(stack, topNode.Right)
		}

		if topNode.Left != nil {
			stack = append(stack, topNode.Left)
		}
	}

	return result
}

//二叉树中序遍历-迭代-左中右
func BinartTreeMiddle(node *TreeNode) []int {
	var result []int
	if node == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		topNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, topNode.Value)
		if topNode.Right != nil {
			node = topNode.Right
		}

	}
	return result
}

//二叉树后序遍历-迭代-左右中
func BinartTreeBack(node *TreeNode) []int {
	var result []int
	if node == nil {
		return result
	}

	//中右左
	stack := make([]*TreeNode, 0)
	stack = append(stack, node)
	for len(stack) > 0 {
		topNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, topNode.Value)

		if topNode.Left != nil {
			stack = append(stack, topNode.Left)
		}

		if topNode.Right != nil {
			stack = append(stack, topNode.Right)
		}
	}

	//左右中
	Reverse(result)
	return result
}

func Reverse(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
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

//递归法解决树
func RecursiveBinartTreePre(node *TreeNode) {
	if node == nil {
		return
	}

	print(node.Value, "-")
	RecursiveBinartTreePre(node.Left)
	RecursiveBinartTreePre(node.Right)
}

func RecursiveBinartTreeMiddle(node *TreeNode) {
	if node == nil {
		return
	}

	RecursiveBinartTreeMiddle(node.Left)
	print(node.Value, "-")
	RecursiveBinartTreeMiddle(node.Right)
}

func RecursiveBinartTreeBack(node *TreeNode) {
	if node == nil {
		return
	}

	RecursiveBinartTreeBack(node.Left)
	RecursiveBinartTreeBack(node.Right)
	print(node.Value, "-")
}

var result [][]int

func RecursiveBinartTreeLevel(node *TreeNode) {
	if node == nil {
		return
	}
	RecursiveBinartTreeLevelOrder(node, 0)

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			print(result[i][j], "-")
		}
	}
}

func RecursiveBinartTreeLevelOrder(node *TreeNode, depth int) {
	if node == nil {
		return
	}

	if len(result) == depth {
		result = append(result, []int{})
	}

	result[depth] = append(result[depth], node.Value)
	RecursiveBinartTreeLevelOrder(node.Left, depth+1)
	RecursiveBinartTreeLevelOrder(node.Right, depth+1)
}
