package leetcode

import "testing"

func TestGroupAnagrams(t *testing.T) {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	GroupAnagrams(str)
}

func TestMaxSubArray(t *testing.T) {
	t.Log(MaxSubArray_2([]int{5, 4, -1, 7, 8}))
}
