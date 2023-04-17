package tree

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
