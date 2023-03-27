package tree

import "testing"

func TestFindInitBinaryTree(t *testing.T) {
	data := []int{2, 4, 5, 1, 3}
	root := InitBinaryTree(data)
	t.Log(root)
}
