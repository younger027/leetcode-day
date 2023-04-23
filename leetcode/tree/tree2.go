package tree

import (
	"fmt"
	"math"
)

//222-完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftNodeNum := countNodes(root.Left)
	rightNodeNum := countNodes(root.Right)

	return leftNodeNum + rightNodeNum + 1
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nodeArray := make([]*TreeNode, 0)
	nodeArray = append(nodeArray, root)
	number := 0

	for len(nodeArray) > 0 {
		size := len(nodeArray)
		number += size
		for i := 0; i < size; i++ {
			node := nodeArray[0]
			nodeArray = nodeArray[1:]

			if node.Left != nil {
				nodeArray = append(nodeArray, node.Left)
			}

			if node.Right != nil {
				nodeArray = append(nodeArray, node.Right)
			}
		}
	}

	return number
}

//110.平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if Isvalid(root) == -1 {
		return false
	}

	return true
}

var result bool

func Isvalid(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftLen := Isvalid(root.Left)
	rightLen := Isvalid(root.Right)

	if leftLen == -1 || rightLen == -1 {
		return -1
	}

	gap := leftLen - rightLen
	if -1 <= gap && gap <= 1 {
		return Max(leftLen, rightLen) + 1
	}

	return -1

}

/*
98. 验证二叉搜索树
中等

给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var pre = math.MinInt32

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isValidBST(root.Left) {
		return false
	}

	if root.Val <= pre {
		return false
	}

	pre = root.Val

	return isValidBST(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return maxDepthPer(root)
}

func maxDepthPer(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return Max(maxDepthPer(root.Left), maxDepth(root.Right)) + 1
}

/*
 preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
 [3,9,20,null,null,15,7]
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder), inorder, 0, len(inorder))
}

func build(preorder []int, ps, pe int, inorder []int, is, ie int) *TreeNode {
	if ps == pe {
		return nil
	}

	root := &TreeNode{
		Val: preorder[ps],
	}
	gap := 0
	for i := is; i < ie; i++ {
		if preorder[ps] == inorder[i] {
			gap = i
			break
		}
	}

	leftLen := gap - is
	root.Left = build(preorder, ps+1, ps+leftLen+1, inorder, is, gap)
	root.Right = build(preorder, ps+leftLen+1, pe, inorder, gap+1, ie)

	return root
}

/*
114. 二叉树展开为链表
中等
1.4K
给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

示例 1：
输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]
*/

func PreT(root *TreeNode) {
	if root == nil {
		return
	}

	result := make([]*TreeNode, 0)
	allResult := make([]*TreeNode, 0)
	result = append(result, root)

	for len(result) > 0 {
		node := result[len(result)-1]
		result = result[0 : len(result)-1]

		fmt.Println("---", node.Val)
		allResult = append(allResult, node)

		if node.Right != nil {
			result = append(result, node.Right)
		}

		if node.Left != nil {
			result = append(result, node.Left)
		}

	}

	for i := 0; i < len(allResult); i++ {
		allResult[i].Left = nil
		if i+1 < len(allResult) {
			allResult[i].Right = allResult[i+1]
		} else {
			allResult[i].Right = nil
		}
	}
}
