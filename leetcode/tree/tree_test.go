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
	data := []int{0}
	root := InitBinaryTree(data, 0)
	//LevelTraverse(root)

	//t.Log(PreorderTraverse(root))
	//t.Log(rightSideView(root))

	//t.Log(MiddleTraverseRecursive(root))
	//t.Log(MiddleTraverse(root))

	//t.Log(PostTraverseRecursive(root))
	//t.Log(PostTraverseTraverse(root))

	//t.Log(largestVals(root))

	//invertTree2(root)
	//LevelTraverse(root)

	//t.Log(isSymmetricFor(root))
	//t.Log(minDepth2(root))

	//t.Log(countNodes(root))
	//t.Log(countNodes2(root))

	//t.Log(isValidBST(root))
	//t.Log(maxDepth(root))

	//PreT(root)
	//t.Log(LevelTraverse(root))

	t.Log(maxPathSum(root))
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)
	LevelTraverse(root)

	data := []int{7, 1, 5, 3, 6, 4}
	t.Log(maxProfit(data))
}
