package leetcode

/*
46. 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：
输入：nums = [1]
输出：[[1]]

提示：
1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/

func Permute(nums []int) [][]int {
	use := make(map[int]bool)
	var result [][]int
	var path []int

	permuteDFS(nums, 0, &result, &path, use)

	return result

}

func permuteDFS(nums []int, index int, result *[][]int, path *[]int, use map[int]bool) {
	if len(*path) == len(nums) {
		dst := make([]int, len(*path))
		copy(dst, *path)
		*result = append(*result, dst)
		return
	}

	for i := 0; i < len(nums); i++ {

		if !use[nums[i]] {
			use[nums[i]] = true
			*path = append(*path, nums[i])
			permuteDFS(nums, i+1, result, path, use)
			*path = (*path)[:len(*path)-1]
			use[nums[i]] = false
		}
	}

}
