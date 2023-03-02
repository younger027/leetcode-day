package leetcode

import "sort"

/*
56. 合并区间
中等

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，
并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/

func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	result := make([][]int, len(intervals))
	index := -1

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		if index == -1 || result[index][1] < intervals[i][0] {
			index += 1
			result[index] = intervals[i]
		} else {
			result[index][1] = GetMax(intervals[i][1], result[index][1])
		}
	}

	return result[:index+1]
}
