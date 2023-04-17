package tree

import "testing"

func TestFindInitBinaryTree(t *testing.T) {
	data := []int{3, 9, 20, -1, -1, 15, 7}
	root := InitBinaryTree(data, 0)
	t.Log(root)

	t.Log(levelOrderBottom(root))

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
	data := []int{1, 2, 3, 4, 5, 5, 5}
	root := InitBinaryTree(data, 0)
	LevelTraverse(root)

	//t.Log(PreorderTraverse(root))
	//t.Log(rightSideView(root))

	//t.Log(MiddleTraverseRecursive(root))
	//t.Log(MiddleTraverse(root))

	//t.Log(PostTraverseRecursive(root))
	//t.Log(PostTraverseTraverse(root))

	//t.Log(largestValues(root))

	//invertTree2(root)
	//LevelTraverse(root)

	//t.Log(isSymmetricFor(root))
	//t.Log(minDepth2(root))

	t.Log(countNodes(root))
	t.Log(countNodes2(root))
}
