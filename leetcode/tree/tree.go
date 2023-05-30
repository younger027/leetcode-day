package tree

import (
	list2 "container/list"
	"context"
	"fmt"
	leetcode "leetcode/leetcode/hot-100"
	"math"
	"sync"
)

type TreeNode struct {
	Val   int
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
		Val:   data[i],
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
		level := fmt.Sprintf("level %d, Val:", i)
		levelResult := make([]int, 0)
		for _, tNode := range nodeArray {
			level += fmt.Sprintf("%d", tNode.Val)
			levelResult = append(levelResult, tNode.Val)
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
		level := fmt.Sprintf("level %d, Val:", i)
		levelResult := make([]int, 0)
		for _, tNode := range nodeArray {
			level += fmt.Sprintf("%d", tNode.Val)
			levelResult = append(levelResult, tNode.Val)
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
		result = append(result, nodeArray[len(nodeArray)-1].Val)

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

	result = append(result, root.Val)
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
		result = append(result, node.Val)
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
	result = append(result, root.Val)
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
			fmt.Println("node---", root.Val)
			result = append(result, root.Val)
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
	result = append(result, root.Val)

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
		result = append(result, root.Val)
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

			levelTotal += tail.Val
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
func largestVals(root *TreeNode) []int {
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
			if node.Val > max {
				max = node.Val
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
	} else if left.Val != right.Val {
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
		} else if left.Val != right.Val {
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

/*
121. 买卖股票的最佳时机
简单
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

示例 1：
输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
示例 2：

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
*/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	low := math.MaxInt32
	result := 0
	for i := 0; i < len(prices); i++ {
		low = Min(low, prices[i])
		result = Max(result, prices[i]-low)
	}

	return result
}

/*
124. 二叉树中的最大路径和
困难
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。
*/

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)

		innerMaxSum := left + root.Val + right
		maxSum = max(maxSum, innerMaxSum)
		outputMaxSum := root.Val + max(left, right) // left,right都是非负的，就不用和0比较了
		return max(outputMaxSum, 0)
	}

	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PreOrderFor(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	list := list2.New()
	var result []int
	list.PushBack(root)

	for list.Len() > 0 {
		node := list.Remove(list.Back()).(*TreeNode)
		result = append(result, node.Val)

		if node.Right != nil {
			list.PushBack(node.Right)
		}

		if node.Left != nil {
			list.PushBack(node.Left)
		}
	}

	return result
}

func MiddleOrderFor(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)

	list := list2.New()

	for list.Len() > 0 || root != nil {
		if root != nil {
			list.PushBack(root)
			root = root.Left
		} else {
			root = list.Remove(list.Back()).(*TreeNode)
			//fmt.Println("node---", root.Val)
			result = append(result, root.Val)
			root = root.Right
		}
	}

	return result
}

func BackOrderFor(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)

	list := list2.New()
	list.PushBack(root)

	for list.Len() > 0 {
		node := list.Remove(list.Back()).(*TreeNode)
		result = append(result, node.Val)

		if node.Left != nil {
			list.PushBack(node.Left)
		}

		if node.Right != nil {
			list.PushBack(node.Right)
		}
	}

	reverse(result)
	return result
}

//a-1234567890, b-abcdc
func WriteA() {
	a := []byte{'1', '2'}
	b := []byte{'a', 'b'}

	ch := make(chan byte)

	var wg sync.WaitGroup
	wg.Add(2)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for i := 0; i < len(a); i++ {
			ch <- a[i]
			ch <- b[i]
		}
		wg.Done()
		cancel()

	}()

	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case s := <-ch:
				//make something-BackOrderFor(nil)
				println("read bytes, ", s)
			case <-ctx.Done():
				close(ch)
				return
			}
		}

	}(ctx)

	wg.Wait()
}

func partition(list []int, low, high int) int {
	pivot := list[low]
	for low < high {
		for low < high && pivot <= list[high] {
			high--
		}
		list[low] = list[high]

		for low < high && pivot >= list[low] {
			low++
		}

		list[high] = list[low]
	}
	//pivot 填补 low位置的空值
	list[low] = pivot
	return low
}
