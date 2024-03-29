package leetcode

import (
	"fmt"
	"sort"
)

/*
75. 颜色分类
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
必须在不使用库内置的 sort 函数的情况下解决这个问题。

示例 1：
输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
示例 2：

输入：nums = [2,0,1]
输出：[0,1,2]

*/

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	//[0, zero]
	//[zero, i]
	//[two, len-1]

	zero := 0
	two := len(nums)

	i := 0
	for i < two {
		if nums[i] == 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			i++
			zero++
		} else if nums[i] == 1 {
			i++
		} else {
			two--
			nums[i], nums[two] = nums[two], nums[i]
		}
	}

	return
}

/*
78. 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]

提示：
1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同
*/

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	path = make([]int, 0)

	BackTrack(nums, &path, &result, 0)
	return result
}

func BackTrack(nums []int, path *[]int, result *[][]int, start int) {
	dst := make([]int, len(*path))
	copy(dst, *path)
	*result = append(*result, *path)

	for i := start; i < len(nums); i++ {
		*path = append(*path, nums[i])
		BackTrack(nums, path, result, i+1)
		*path = (*path)[0 : len(*path)-1]
	}
}

//定义全局变量，不采用指针传递的方式
var (
	result [][]int
	path1  []int
)

func subsetss(nums []int) [][]int {
	result = make([][]int, 0)
	path1 = make([]int, 0)

	BackTracks(nums, 0)
	return result
}

func BackTracks(nums []int, start int) {
	dst := make([]int, len(path1))
	copy(dst, path1)
	fmt.Println("dst:", dst)

	result = append(result, dst)

	for i := start; i < len(nums); i++ {
		path1 = append(path1, nums[i])
		BackTracks(nums, i+1)
		path1 = path1[0 : len(path1)-1]
	}

	fmt.Println("result:", result)
}

/*
90. 子集 II
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

示例 1：
输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
示例 2：
输入：nums = [0]
输出：[[],[0]]

提示：
1 <= nums.length <= 10
-10 <= nums[i] <= 10*/

func subsetsWithDup(nums []int) [][]int {
	result = make([][]int, 0)
	path = make([]int, 0)

	sort.Ints(nums)
	subsetsWithDupDFS(nums, 0)

	return result
}

func subsetsWithDupDFS(nums []int, start int) {
	//回溯法
	//结束条件，i-1, i相同
	dst := make([]int, len(path))
	copy(dst, path)
	fmt.Println("dst:", dst)

	result = append(result, dst)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		path = append(path, nums[i])
		subsetsWithDupDFS(nums, i+1)
		path = path[0 : len(path)-1]

	}
}

/*
96. 不同的二叉搜索树
中等
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
dp五部曲:
	1.状态定义:dp[i]为当有i个节点时,一共可以组成的二叉搜索树数目
	2.状态转移:dp[3]=dp[0]*dp[2]+dp[1]*dp[1]+dp[2]*dp[0]
		可以比喻成前面一项是左子树情况数,后面一项是右子树情况数,相加即可
		即:dp[i]=∑dp[j]*dp[i-1-j],其中j∈[0,i-1]
	3.初始化:dp[0]=1,dp[1]=dp[0]*dp[0]=1,dp[2]=dp[0]*dp[1] + dp[1]*dp[0]
	4.遍历顺序:正序
	5.返回形式:返回dp[n]
*/

func numTrees(n int) int {
	dpt := make([]int, n+1)

	dpt[0] = 1

	for i := 1; i <= n; i++ {
		fmt.Println("------", i)

		for j := 0; j < i; j++ {
			dpt[i] += dpt[j] * dpt[i-1-j]
			fmt.Println(j, i-1-j)
		}
	}

	return dpt[n]
}
