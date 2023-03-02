package leetcode

/*
55. 跳跃游戏
中等
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。

示例 1：
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

示例 2：
输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。

提示：
1 <= nums.length <= 3 * 104
0 <= nums[i] <= 105
*/

//CanJump solution 1 by self.exec time out
func CanJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return true
	}

	return jump(nums, 0, nums[0])
}

func jump(nums []int, start, limit int) bool {
	if start+limit >= len(nums)-1 {
		return true
	}

	flag := false
	for i := 1; i <= limit; i++ {
		flag = flag || jump(nums, start+i, nums[start+i])
		if flag {
			return true
		}
	}

	return flag
}

//solution 2
//k means the longest index we can arrive

func CanJump_2(nums []int) bool {
	k := 0
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}

		k = GetMax(k, i+nums[i])
	}

	return true
}
