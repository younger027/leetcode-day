package tree

import (
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

//前序遍历
func preorderTraverse(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	result = append(result, root.Value)
	result = append(result, preorderTraverse(root.Left)...)
	result = append(result, preorderTraverse(root.Right)...)

	return result
}

func InitBinaryTree(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var i int
	root := &TreeNode{Value: preorder[0]}
	for i = 1; i < len(preorder); i++ {
		if preorder[i] > root.Value {
			break
		}
	}
	root.Left = InitBinaryTree(preorder[1:i])
	root.Right = InitBinaryTree(preorder[i:])
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

var number int

func DiameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := DiameterOfBinaryTree(root.Left)
	right := DiameterOfBinaryTree(root.Right)

	max := left + right
	number = leetcode.Max(number, max)

	return 1 + leetcode.Max(left, right)
}

//层序遍历
func LevelTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	nodeArray := make([]*TreeNode, 0)
	nodeArray = append(nodeArray, root)
	i := 0
	for len(nodeArray) > 0 {
		level := fmt.Sprintf("level %d, value:", i)
		for _, tNode := range nodeArray {
			level += fmt.Sprintf("%d", tNode.Value)
			node := nodeArray[0]
			nodeArray = nodeArray[1:]
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
}
