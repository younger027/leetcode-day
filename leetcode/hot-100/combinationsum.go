package leetcode

import "sort"

/*
39. 组合总和
给你一个 无重复元素 的整数数组candidates和一个目标整数 target ，找出candidates中可以使数字和为目标数target的所有不同组合，并以列表形式返回。
你可以按 任意顺序 返回这些组合。candidates中的 同一个数字可以无限制重复被选取。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为target的不同组合数少于 150 个。

示例 1：
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。

示例 2：
输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]

示例 3：
输入: candidates = [2], target = 1
输出: []

提示：

1 <= candidates.length <= 30
2 <= candidates[i] <= 40
candidates 的所有元素 互不相同
1 <= target <= 40

*/

func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var path []int

	if len(candidates) == 0 {
		return result
	}

	//sort for cut branch
	sort.Ints(candidates)

	dfsCombination(candidates, 0, target, &result, &path)

	return result
}

func dfsCombination(candidates []int, begin, target int, result *[][]int, path *[]int) {
	if target == 0 {
		dst := make([]int, len(*path))
		copy(dst, *path)
		*result = append(*result, dst)
		return
	}

	for i := begin; i < len(candidates); i++ {
		if target < 0 {
			//because sort already,so target < 0 do not need for loop other number
			break
		}

		*path = append(*path, candidates[i])
		dfsCombination(candidates, i, target-candidates[i], result, path)
		*path = (*path)[:len(*path)-1]
	}

}