package interview

import "testing"

func TestInitTreeNode(t *testing.T) {
	node := InitTreeNode([]int{1, 2, 3, 4, 5, 6, 7}, 0)
	t.Log(node)

	t.Log(BinartTreePre(node))
	t.Log(BinartTreeMiddle(node))
	t.Log(BinartTreeBack(node))
	t.Log(BinartTreeLevel(node))

	RecursiveBinartTreePre(node)
	t.Log("\n")
	RecursiveBinartTreeMiddle(node)
	t.Log("\n")
	RecursiveBinartTreeBack(node)
	t.Log("\n")
	RecursiveBinartTreeLevel(node)
}
