package leetcode

import "testing"

func TestGroupAnagrams(t *testing.T) {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	GroupAnagrams(str)
}

func TestMaxSubArray(t *testing.T) {
	t.Log(MaxSubArray_2([]int{5, 4, -1, 7, 8}))
}

func TestCanJump(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	t.Log(CanJump(nums))

	t.Log(CanJump_2(nums))
}

func TestMerge(t *testing.T) {
	nums := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//nums := [][]int{{1, 4}, {4, 5}}
	//nums := [][]int{{1, 4}, {0, 4}}
	t.Log(Merge(nums))
}

func TestUniquePaths(t *testing.T) {
	t.Log(UniquePaths(3, 2))
}

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}

	t.Log(MaxSlidingWindow(nums, 3))
}

func TestLongestPalindrome2(t *testing.T) {

	t.Log(LongestPalindrome2("abbac"))
}
