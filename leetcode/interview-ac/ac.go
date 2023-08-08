package interview

import (
	"bytes"
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

var (
	path       []int
	resultData [][]int
)

//1155. 掷骰子等于目标和的方法数
//可以通过测试用例，会超时
func numRollsToTargetSelf(n int, k int, target int) int {
	//resultData = make([][]int, 0)
	Trace(n, k, target, 0)

	return len(resultData)
}

func Trace(n int, k int, target int, current int) {
	if current > target {
		return
	}

	if current == target && len(path) == n {
		tmp := make([]int, len(path))
		copy(tmp, path)
		resultData = append(resultData, tmp)
		return
	}

	for j := 1; j <= k; j++ {
		path = append(path, j)
		current += j
		Trace(n, k, target, current)
		path = path[:len(path)-1]
		current -= j
	}

}

func Sum(path []int) int {
	sum := 0
	for _, item := range path {
		sum += item
	}

	return sum
}

/*dp[i][j]代表i个骰子凑成target=j的方案数
dp[i][j] +== dp[i-1][j-[1~k]].
第i个骰子的数字是1，当骰子是1时，那么dp[i-1][j-1]就代表i-1个骰子骰出j-1的种类有多少。
第i个骰子的数字是2，当骰子是2时，那么dp[i-1][j-2]就代表i-1个骰子骰出j-2的种类有多少。
一直到k。思路主要是反着来的。最后一颗骰子的范围在1~k,那么当第i颗投出这个结果时，种类数就依赖
前i-1颗能投出j-k的数量了。慢慢品 你可以的。
初始化：dp[0][j]:0颗骰子投不出其他的j，只能dp[0][0]=1,其他的dp[0][j] = 0,不可能抛出来
遍历顺序：背包问题，先遍历物品，再背包

*/
func numRollsToTarget(n int, k int, target int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, target+1)
	}

	dp[0][0] = 1
	mod := int(1e9 + 7)
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			for z := 1; z <= k && z <= j; z++ {
				dp[i][j] = (dp[i][j] + dp[i-1][j-z]) % mod
			}
		}
	}

	return dp[n][target]
}

/*
# 1768 交替合并字符串
*/
func mergeAlternately(word1 string, word2 string) string {
	len1 := len(word1)
	len2 := len(word2)

	var result bytes.Buffer
	for i, j := 0, 0; i < len1 || j < len2; i, j = i+1, j+1 {
		if i < len1 {
			result.WriteByte(word1[i])
		}
		if j < len2 {
			result.WriteByte(word2[j])
		}
	}

	return result.String()
}

func gcdOfStrings(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	if len1 > len2 {
		str1, str2 = str2, str1
		len1, len2 = len2, len1
	}

	for i := len1; i > 0; i-- {
		for j, k := 0, 0; k < len2; j, k = j+1, k+1 {
			oldJ := j
			j = j % i
			if str1[j] != str2[k] {
				break
			}

			if k == len2-1 && oldJ == i-1 {
				return str1[:i]
			}
		}
	}

	return ""

}
