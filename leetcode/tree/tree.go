package tree

import (
	list2 "container/list"
	"fmt"
	leetcode "leetcode/leetcode/hot-100"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// 记录最大深度
var res int

// 记录遍历到的节点的深度
var depth int

func MaxDepth(root *TreeNode) int {
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	depth += 1
	if root.Left == nil && root.Right == nil {
		res = leetcode.Max(res, depth)
	}
	traverse(root.Left)
	traverse(root.Right)
	depth -= 1
}

func MaxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := MaxDepth2(root.Left)
	rightMax := MaxDepth2(root.Right)

	return leetcode.Max(leftMax, rightMax) + 1
}

func InitBinaryTree(data []int, i int) *TreeNode {
	if data[i] == -1 {
		return nil
	}

	root := &TreeNode{
		Value: data[i],
		Left:  nil,
		Right: nil,
	}

	if i < len(data) && 2*i+1 < len(data) {
		root.Left = InitBinaryTree(data, 2*i+1)
	}

	if i < len(data) && 2*i+2 < len(data) {
		root.Right = InitBinaryTree(data, 2*i+2)
	}

	return root
}

/*
543. 二叉树的直径
简单
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。
*/

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := 0
	Depth(root, &max)
	fmt.Println("max ", max)
	return max
}

func Depth(root *TreeNode, maxDepth *int) int {
	if root == nil {
		return 0
	}

	left := Depth(root.Left, maxDepth)
	right := Depth(root.Right, maxDepth)

	*maxDepth = Max(*maxDepth, left+right)

	return 1 + Max(left, right)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//层序遍历
func LevelTraverse(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	nodeArray := make([]*TreeNode, 0)
	nodeArray = append(nodeArray, root)
	i := 0
	for len(nodeArray) > 0 {
		level := fmt.Sprintf("level %d, value:", i)
		levelResult := make([]int, 0)
		for _, tNode := range nodeArray {
			level += fmt.Sprintf("%d", tNode.Value)
			levelResult = append(levelResult, tNode.Value)
			node := nodeArray[0]
			nodeArray = nodeArray[1:]
			if node.Left != nil {
				nodeArray = append(nodeArray, node.Left)
			}

			if node.Right != nil {
				nodeArray = append(nodeArray, node.Right)
			}
		}

		result = append(result, levelResult)
		fmt.Println(level)
		i += 1
	}

	return result
}

//层序遍历二
func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	nodeArray := make([]*TreeNode, 0)
	nodeArray = append(nodeArray, root)
	i := 0
	for len(nodeArray) > 0 {
		level := fmt.Sprintf("level %d, value:", i)
		levelResult := make([]int, 0)
		for _, tNode := range nodeArray {
			level += fmt.Sprintf("%d", tNode.Value)
			levelResult = append(levelResult, tNode.Value)
			node := nodeArray[0]
			nodeArray = nodeArray[1:]
			if node.Left != nil {
				nodeArray = append(nodeArray, node.Left)
			}

			if node.Right != nil {
				nodeArray = append(nodeArray, node.Right)
			}
		}

		result = append(result, levelResult)
		fmt.Println(level)
		i += 1
	}

	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}

	return result
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	nodeArray := make([]*TreeNode, 0)
	nodeArray = append(nodeArray, root)
	i := 0
	for len(nodeArray) > 0 {
		level := fmt.Sprintf("level %d, value:", i)
		for index, tNode := range nodeArray {
			level += fmt.Sprintf("%d", tNode.Value)

			node := nodeArray[0]
			nodeArray = nodeArray[1:]
			if index == len(nodeArray)-1 {
				result = append(result, tNode.Value)
			}
			
			if node.Left != nil {
				nodeArray = append(nodeArray, node.Left)
			}

			if node.Right != nil {
				nodeArray = append(nodeArray, node.Right)
			}
		}

		fmt.Println(level)
		i += 1
	}

	return result
}

//前序遍历-递归解法
func PreorderTraverseRecursive(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	result = append(result, root.Value)
	result = append(result, PreorderTraverseRecursive(root.Left)...)
	result = append(result, PreorderTraverseRecursive(root.Right)...)

	return result
}

func PreorderTraverse(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	list := list2.New()
	list.PushBack(root)

	for list.Len() > 0 {
		node := list.Remove(list.Back()).(*TreeNode)
		result = append(result, node.Value)
		if node.Right != nil {
			list.PushBack(node.Right)
		}

		if node.Left != nil {
			list.PushBack(node.Left)
		}
	}

	return result
}
