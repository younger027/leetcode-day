package tree

import "testing"

func TestFindInitBinaryTree(t *testing.T) {
	data := []int{3, 9, 20, -1, -1, 15, 7}
	root := InitBinaryTree(data, 0)
	t.Log(root)

	t.Log(levelOrderBottom(root))

	t.Log(rightSideView(root))
	LevelTraverse(root)

	t.Log("PreorderTraverseRecursive-", PreorderTraverseRecursive(root))
}

func TestDiameterOfBinaryTree(t *testing.T) {
	data := []int{1, 2}
	root := InitBinaryTree(data, 0)
	LevelTraverse(root)

	t.Log(diameterOfBinaryTree(root))

}

func TestPreorderTraverse(t *testing.T) {
	data := []int{3, 9, 20, -1, -1, 15, 7}
	root := InitBinaryTree(data, 0)
	LevelTraverse(root)

	t.Log(PreorderTraverse(root))

}
