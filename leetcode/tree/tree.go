package tree

import (
	list2 "container/list"
	"fmt"
	leetcode "leetcode/leetcode/hot-100"
	"math"
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

//111-二叉树的最小深度,递归法
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftLen := minDepth(root.Left)
	rightLen := minDepth(root.Right)

	if root.Left == nil {
		return 1 + rightLen
	}

	if root.Right == nil {
		return 1 + leftLen
	}

	return 1 + Min(leftLen, rightLen)
}

//迭代法
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nodeArray := make([]*TreeNode, 0)
	nodeArray = append(nodeArray, root)

	depthParams := 0
	for len(nodeArray) > 0 {
		size := len(nodeArray)
		depthParams += 1
		for i := 0; i < size; i++ {
			node := nodeArray[0]
			nodeArray = nodeArray[1:]

			if node.Left != nil {
				nodeArray = append(nodeArray, node.Left)
			}

			if node.Right != nil {
				nodeArray = append(nodeArray, node.Right)
			}

			if node.Left == nil && node.Right == nil {
				return depthParams
			}
		}
	}

	return depthParams
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

func Min(a, b int) int {
	if a < b {
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
		result = append(result, nodeArray[len(nodeArray)-1].Value)

		for i := range nodeArray {
			fmt.Println(i)
			node := nodeArray[0]
			nodeArray = nodeArray[1:]

			if node.Left != nil {
				nodeArray = append(nodeArray, node.Left)
			}

			if node.Right != nil {
				nodeArray = append(nodeArray, node.Right)
			}
		}

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

//中序遍历的递归写法
func MiddleTraverseRecursive(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	result = append(result, MiddleTraverseRecursive(root.Left)...)
	result = append(result, root.Value)
	result = append(result, MiddleTraverseRecursive(root.Right)...)

	return result
}

func MiddleTraverse(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := list2.New()
	for stack.Len() > 0 || root != nil {
		if root != nil {
			stack.PushBack(root)
			fmt.Println("PushBack node---", stack.Len())

			root = root.Left
		} else {
			root = stack.Remove(stack.Back()).(*TreeNode)
			fmt.Println("node---", root.Value)
			result = append(result, root.Value)
			root = root.Right

		}
	}

	return result
}

//后序遍历的递归写法
func PostTraverseRecursive(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	result = append(result, PostTraverseRecursive(root.Left)...)
	result = append(result, PostTraverseRecursive(root.Right)...)
	result = append(result, root.Value)

	return result
}

func PostTraverseTraverse(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := list2.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		root = stack.Remove(stack.Back()).(*TreeNode)
		result = append(result, root.Value)
		if root.Left != nil {
			stack.PushBack(root.Left)
		}

		if root.Right != nil {
			stack.PushBack(root.Right)
		}
	}

	reverse(result)
	return result
}

func reverse(data []int) {
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-i-1] = data[len(data)-i-1], data[i]
	}
}

//637.二叉树的层平均值
func averageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)
	if root == nil {
		return result
	}

	levelArray := make([]*TreeNode, 0)
	levelArray = append(levelArray, root)
	for len(levelArray) > 0 {
		levelTotal := 0
		sz := len(levelArray)
		for i := 0; i < sz; i++ {
			tail := levelArray[0]
			levelArray = levelArray[1:]

			levelTotal += tail.Value
			if tail.Left != nil {
				levelArray = append(levelArray, tail.Left)
			}

			if tail.Right != nil {
				levelArray = append(levelArray, tail.Right)
			}
		}

		result = append(result, float64(levelTotal)/float64(sz))
	}

	return result
}

//429.N叉树的层序遍历
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	levelArray := make([]*Node, 0)
	levelArray = append(levelArray, root)
	for len(levelArray) > 0 {
		var tmp []int
		sz := len(levelArray)
		for i := 0; i < sz; i++ {
			node := levelArray[0]
			levelArray = levelArray[1:]
			tmp = append(tmp, node.Val)

			for _, v := range node.Children {
				levelArray = append(levelArray, v)
			}
		}
		result = append(result, tmp)
	}

	return result
}

//515. 在每个树行中找最大值
func largestValues(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	levelArray := make([]*TreeNode, 0)
	levelArray = append(levelArray, root)
	for len(levelArray) > 0 {
		sz := len(levelArray)
		max := math.MinInt64
		for i := 0; i < sz; i++ {
			node := levelArray[0]
			levelArray = levelArray[1:]
			if node.Value > max {
				max = node.Value
			}
			if node.Left != nil {
				levelArray = append(levelArray, node.Left)
			}
			if node.Right != nil {
				levelArray = append(levelArray, node.Right)

			}
		}

		result = append(result, max)
	}

	return result
}

type NewNode struct {
	Val   int
	Left  *NewNode
	Right *NewNode
	Next  *NewNode
}

//116. 填充每个节点的下一个右侧节点指针
func connect(root *NewNode) *NewNode {
	if root == nil {
		return root
	}

	levelArray := make([]*NewNode, 0)
	levelArray = append(levelArray, root)

	for len(levelArray) > 0 {
		sz := len(levelArray)
		for i := 0; i < sz; i++ {
			node := levelArray[0]
			levelArray = levelArray[1:]
			if node.Left != nil {
				levelArray = append(levelArray, node.Left)
			}
			if node.Right != nil {
				levelArray = append(levelArray, node.Right)

			}

			if i < sz-1 {
				node.Next = levelArray[0]
			}

		}
	}

	return root
}

//226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	stack := list2.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			stack.PushBack(node.Left)
		}

		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}

	return root
}

//LeetCode：101. 对称二叉树
func compare(left, right *TreeNode) bool {
	//终止条件, // 首先排除空节点的情况
	if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left == nil && right == nil {
		return true
	} else if left.Value != right.Value {
		// 排除了空节点，再排除数值不相同的情况
		return false
	}

	// 此时就是：左右节点都不为空，且数值相同的情况
	// 此时才做递归，做下一层的判断
	outer := compare(left.Left, right.Right)
	inner := compare(left.Right, right.Left)

	return outer && inner
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

//迭代法

func isSymmetricFor(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root.Left)
	stack = append(stack, root.Right)

	for len(stack) > 0 {
		left := stack[0]
		stack = stack[1:]
		right := stack[0]
		stack = stack[1:]

		if left == nil && right == nil {
			continue
		}

		if left == nil && right != nil {
			return false
		} else if left != nil && right == nil {
			return false
		} else if left.Value != right.Value {
			// 排除了空节点，再排除数值不相同的情况
			return false
		}

		stack = append(stack, left.Left, right.Right)
		stack = append(stack, left.Right, right.Left)

	}

	return true
}

//572. 另一棵树的子树
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return compare(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
