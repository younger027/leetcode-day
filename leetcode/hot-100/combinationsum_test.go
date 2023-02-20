package leetcode

import "testing"

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 5}
	target := 8

	t.Log(CombinationSum(candidates, target))
}
