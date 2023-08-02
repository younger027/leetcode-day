package interview

import "testing"

func TestInitTreeNode(t *testing.T) {
	node := InitTreeNode([]int{1, 2, 3, 4, 5, 6, 7}, 0)
	t.Log(node)
}
