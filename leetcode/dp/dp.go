package dp

import (
	"fmt"
	leetcode "leetcode/leetcode/hot-100"
	"math"
	"sort"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)

	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}

	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
				fmt.Println(i, j, dp[i][j])
			}

		}
	}

	return dp[m-1][n-1]
}

func integerBreak(n int) int {
	//dp[i]代表，i的拆分后的最大乘积和
	//状态转移方程：dp[i] = dp[i-j]*0-->j这个范围的最大值那就是j了、得：dp[i-j]*j，还有一个是(i-j)*j,取最大值即可
	dp := make([]int, n+1)
	dp[2] = 1

	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = leetcode.Max(dp[i], leetcode.Max(dp[i-j]*j, (i-j)*j))
		}
	}

	for i := 3; i <= n; i++ {
		fmt.Println(dp[i], i)
	}
	return dp[n]
}

func numTrees(n int) int {
	//状态转移方程式：dp[i]+=dp[i]*dp[i-1-j]
	dp := make([]int, n+1)

	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[i-j] * dp[j-1]
		}
	}

	return dp[n]
}

// 416-分割等和子集
func canPartition(nums []int) bool {
	sum := 0
	dp := make([]int, 10001)
	dp[0] = 0
	for _, n := range nums {
		sum += n
	}

	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = leetcode.Max(dp[j], dp[j-nums[i]]+nums[i])
		}

	}

	if dp[target] == target {
		return true
	}

	return false
}

// 1049. 最后一块石头的重量 II
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, n := range stones {
		sum += n
	}

	target := sum / 2
	dp := make([]int, target+1)

	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = leetcode.Max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	//最后dp[target]里是容量为target的背包所能背的最大重量。
	//那么分成两堆石头，一堆石头的总重量是dp[target]，另一堆就是sum - dp[target]。
	return sum - dp[target] - dp[target]
}

func isEqueal(nums, used []int) bool {
	l, r := 0, 0
	for i, n := range used {
		if n == 0 {
			l += nums[i]
		} else {
			r += nums[i]
		}
	}

	return l == r
}

func findRepeatNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i += 1
			continue
		}

		if nums[nums[i]] == nums[i] {
			return nums[i]
		}

		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}

	return 0
}

func findTargetSumWays(nums []int, target int) int {
	//假设加法的总和为x，那么减法对应的总和就是sum - x。
	//所以我们要求的是 x - (sum - x) = target
	//x = (target + sum) / 2
	sum := 0

	for _, i := range nums {
		sum += i
	}

	if int(math.Abs(float64(target))) > sum {
		return 0
	}

	if (sum+target)%2 == 1 {
		return 0
	}
	bag := (sum + target) / 2

	dp := make([]int, bag+1)

	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[bag]
}

// 动态规划题目，dp[i][j]代表最多有i个0 j个1的strs数组最大子集的大小
// dp[i][j] = max(dp[i][j], dp[i-当前str的0的个数][j-当前str的1的个数]+1)
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			dp[i][j] = 0
		}
	}

	for _, item := range strs {
		countZero, countOne := 0, 0
		for _, c := range item {
			if c == '0' {
				countZero += 1
			} else {
				countOne += 1
			}
		}

		//倒序
		for i := m; i >= countZero; i-- {
			for j := n; j >= countOne; j-- {
				dp[i][j] = leetcode.Max(dp[i][j], dp[i-countZero][j-countOne]+1)
			}
		}
	}

	return dp[m][n]
}

// 动态规划，完全背包问题
// 定义dp dp[i]代表。总金额为i时，可以凑成i的硬币组合数
// dp[i] += dp[i-coins[i]]
// 初始化 dp[0]=0
// 遍历顺序：完全背包的遍历顺序，外层遍历物品，内层从小到大遍历容量
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for i, _ := range coins {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}

// 动态规划：dp[i] 代表target是i时，排列的个数
// dp[i] += dp[i-nums[i]]
// 本题求完全背包的排列数，遍历顺序应该是先从小到大遍历容量，在遍历物品
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	dpp := make([]int, target+1)
	dpp[0] = 1
	//外层遍历
	for i := 0; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i >= nums[j] {
				dpp[i] += dpp[i-nums[j]]
				fmt.Println("----", j, dpp[j], i)
			}
		}
	}

	for i, _ := range nums {
		for j := nums[i]; j <= target; j++ {
			dp[j] += dp[j-nums[i]]
			fmt.Println("----", dp[j-nums[i]], dp[j], i)
		}

	}

	return dp[target]
}

// 爬楼梯
func climbStairsThird(n int) int {
	if n == 0 {
		return 0
	}

	//定义dp函数，dp[i]代表爬i层阶梯，有dp[i]种排列
	dp := make([]int, n+1)
	dp[0] = 1

	//求组合问题，1，3； 3，1；是两种组合。所以需要先遍历重量，而非物品；
	nums := []int{1, 2}
	for i := 1; i <= n; i++ {
		for j := 0; j < len(nums); j++ {
			dp[i] += dp[i-nums[j]]
		}
	}

	return dp[n]
}

// leetcode 322
// 定义dp函数，dp[i],可以凑成总金额为i的最少得硬币个数是dp[i]
// dp[i] = min(dp[i], dp[i-coins[j]]+1)；消耗掉本次金币后，最小硬币数是dp[i-coins[j]]，
// 然后加1代表本次需要使用的金币
// 初始化：dp[0] = 0 金额为0时，硬币个数总和肯定为0
// 遍历顺序：金币无限次使用，完全背包。需要先遍历重量，让每次物品都能重复使用。
// 如果先遍历物品的话，物品只能使用一次
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && dp[i-coins[j]] != math.MaxInt {
				dp[i] = leetcode.Min(dp[i], dp[i-coins[j]]+1)
			}

		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

// 279. 完全平方数
// 此题和上面的题类似。
func numSquares(n int) int {
	dp := make([]int, +1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i <= n; i++ { //背包
		for j := 1; j*j <= i; j++ { //物品
			dp[i] = leetcode.Min(dp[i], dp[i-j*j]+1)

		}
	}

	if dp[n] == math.MaxInt {
		return -1
	}
	return dp[n]
}

// 139. 单词拆分
// 完全背包：dp[i] 代表0-i的字符串，可以拆分成一个或多个dict中的字串
// dp[i]: if j>i && dp[i]==true && i-j的字串在dict中，那么dp[j] = true
// 初始化问题
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	//wordDictSet := make(map[string]bool)
	//for _, w := range wordDict {
	//	wordDictSet[w] = true
	//}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && leetcode.FindSubStr(s[j:i], wordDict) {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}

// 300. 最长递增子序列
// dp[i]： 以i结尾的最长递增子序列的长度
// if nums[i] > nums[j],  dp[i]= max(dp[i], dp[j]+1)
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)+1)
	for i := range dp {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = leetcode.Max(dp[i], dp[j]+1)
				if dp[i] > res {
					res = dp[i]
				}
			}
		}
	}

	return res
}

// 674. 最长连续递增序列
// dp[i]: 以i结尾的最长连续递增序列是dp[i]
// dp[i]: if i=j+1 && nums[i]>nums[j]; dp[i] = dp[j] + 1
// 初始化全部为1，
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)+1)
	for i := range dp {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			if dp[i] > res {
				res = dp[i]
			}
		}
	}

	return res
}

// 674的贪心算法
func findLengthOfLCIS2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 1
	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count += 1
		} else {
			count = 1
		}
		if count > res {
			res = count
		}
	}

	return res
}

// 718. 最长重复子数组
// dp[i][j]:dp[i][j] ：以下标i - 1为结尾的A，和以下标j - 1为结尾的B，
// 最长重复子数组长度为dp[i][j]（特别注意： “以下标i - 1为结尾的A” 标明一定是 以A[i-1]为结尾的字符串 ）
// if nums1[i-1]== nums[j-1], 因为dp的定义是以i-1, j-1结尾的下标，所以此时的判断条件是nums1[i-1]== nums[j-1]。dp[i][j] = dp[i-1][j-1]+1
func findLength(nums1 []int, nums2 []int) int {
	m := len(nums1) + 1
	n := len(nums2) + 1
	res := 0

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}

			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res
}

// 1143. 最长公共子序列 text1 = "ece", text2 = "abcde"
func longestCommonSubsequenceSelf(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)

	if len1 > len2 {
		text1, text2 = text2, text1
	}

	text1Map := make(map[byte]int, len(text1))
	for i := 0; i < len(text1); i++ {
		text1Map[text1[i]] = i + 1
	}

	text2Map := make(map[byte]int, len(text1))
	for i := 0; i < len(text2); i++ {
		text2Map[text2[i]] = i + 1
	}
	res := 0
	index := 0
	lastIndex := 0
	for i := 0; i < len(text1); i++ {
		newIndex, ok := text2Map[text1[i]]
		if ok {
			if newIndex > lastIndex {
				lastIndex = newIndex
				index += 1
			} else {
				index = 1
			}

		}

		if index > res {
			res = index
		}
	}

	return res
}

// 1143. 最长公共子序列；定义动态规划dp[i][j]代表text1的0，i-1和text2的0，j-1的最长公共子序列是dp[i][j]
// 状态转移方程:dp[i][j]:
// if text1[i-1]==tex2[j-1]; dp[i][j] = dp[i-1][j-1]+1
// if not 那就看看text1[0, i - 2]与text2[0, j - 1]的最长公共子序列和text1[0, i-1]与text2[0, j-2]的最长公共子序列，
// 取最大的。dp[i][j] = max(dp[i][j-1],dp[i-1][j])
// 初始化数组，根据定义dp[0][j]代表text1的空串和0-j-1的text2最长公共子序列，自然是0,dp[i][0]自然也是0
func longestCommonSubsequenceDp(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = leetcode.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

// 1035. 不相交的线 和上面的题一样，意思就是求最长公共子序列
// 动态五部曲：1.定义动态规则dp[i][j],代表的是nums1以i-1结尾的字符串和nums2以j-1结尾的字符串的最长公共子序列是dp[i][j]
// 规则转化：2.当nums1[i-1] == nums[j-1],dp[i][j]= dp[i-1][j-1]+1,
// 不等于时呢：dp[i][j]就是(nums1的0--i-1 和nums2的0-j)和(nums1的0--i 和nums2的0-j-1)哪个子序列最长
// dp[i][j] = max(dp[i][j-1], dp[i-1][j])
// 循环遍历即可，因为定义的是i-1,j-1结尾，那么就不用去初始化dp[i][0], dp[0][j]这些行列，默认是0，也是合理的
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			dp[i][j] = leetcode.Max(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m][n]
}

/*
53. 最大子数组和
dp[i]：代表以i结尾的数组，其最大子数组和是dp[i]
公式：max(dp[i-1]+nums[i], nums[i])
初始化dp[0]=nums[0]
顺序遍历即可
*/

func maxSubArray(nums []int) int {
	m := len(nums)
	dp := make([]int, m+1)
	dp[0] = nums[0]
	res := math.MinInt
	for i := 1; i <= m; i++ {
		dp[i] = leetcode.Max(dp[i-1]+nums[i], nums[i])

		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}

/*
392. 判断子序列
dp[i][j]:s以i-1结尾，t以j-1结尾的的子序列是长度是dp[i][j]
当s[i-1]==t[j-1]时，dp[i][j]=dp[i-1][j-1]+1,否则相当于t删除了t[j-1],那就要对比s[i-1]和t[j-2]的情况。即dp[i][j-1]
初始化的时候，按照dp的定义，dp[0][j]和dp[i][0]都应该是0
*/
func isSubsequence(s string, t string) bool {
	m := len(s)
	n := len(t)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]

			}
		}
	}

	return dp[m][n] == m
}

/*
583. 两个字符串的删除操作
定义dp[i][j],是以i-1结尾的word1和j-1结尾的word2，相同的最小改变
方程：当word1[i-1]==word[j-1],dp[i][j] = dp[i-1][j-1]
不相等时，删除word1的i-1, dp[i-1][j]+1
删除word2的j-1, dp[i][j-1]+1
都删除,dp[i-1][j-1]+2.取三者的最小值就行
***初始化：dp[i][0]的定义，word1要变成空的word2需要删除i-1个元素。那么dp[i][0]=i,dp[0][j]=j
*/

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = leetcode.Min(dp[i][j-1]+1, leetcode.Min(dp[i-1][j-1]+2, dp[i-1][j]+1))
			}
		}
	}

	return dp[m][n]
}

/*
72. 编辑距离
dp[i][j] 表示以下标i-1为结尾的字符串word1，和以下标j-1为结尾的字符串word2，最近编辑距离为dp[i][j]。
if word1[i-1]==word2[j-1]，dp[i][j]= dp[i-1][j-1]
else
增加/删除。word1增加一个dp[i][j]= dp[i][j-1]+1。相当于word2删除了一个。
word2增加一个dp[i][j]= dp[i-1][j]+1。相当于word1删除了一个。
更改：word1的i-1更改成word2的j-1，dp[i][j]= dp[i-1][j-1]+1

*/

func MinDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	dp := make([][]int, l1+1)

	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}

	//列
	for i := 1; i <= l1; i++ {
		//下面任一都可以，第二条比较好理解
		//dp[i][0] = dp[i-1][0] + 1
		dp[i][0] = i
	}

	//行
	for j := 1; j <= l2; j++ {
		//下面任一都可以，第二条比较好理解
		//dp[0][j] = dp[0][j-1] + 1
		dp[0][j] = j
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minNumThree(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[l1][l2]
}

func minNumThree(a, b, c int) int {
	if a > b {
		if c > b {
			return b
		} else {
			return c
		}
	} else {
		if c > a {
			return a
		} else {
			return c
		}
	}
}

/*
订单拆解：一个订单n个商品，假如按重量拆为m个子订单，忽略价格因素，
每个子订单的商品重量不超过x千克，如何分使得子订单数量m最少

n个商品的总重量是一定的，每个子订单的重量也是有限制的。可以采用最简单的方法，排序后，双指针头尾相加，根据和x值的对比，左右++--
*/

var (
	res  = make([][]int, 0)
	path = make([]int, 0)
)

func packageSplit(x int, weight []int) int {
	sort.Ints(weight)
	backTrace(x, 0, 0, weight)
	for _, item := range res {
		for _, ii := range item {
			fmt.Printf("%d-", ii)
		}
		fmt.Println("")
	}
	return len(res)
}

func backTrace(x int, currentWeight, startIndex int, weight []int) {
	if currentWeight > x {
		dst := make([]int, len(path))
		copy(dst, path)
		res = append(res, dst)
		return
	}

	for i := startIndex; i < len(weight); i++ {
		path = append(path, weight[i])
		currentWeight = currentWeight + weight[i]
		backTrace(x, currentWeight, i+1, weight)
		path = path[0 : len(path)-1]
		currentWeight = currentWeight - weight[i]
	}

}

/*
647. 回文子串
dp[i][j]代表以i-1,j-1段是否是回文字串
推导公式：
当s[i]不等于s[j]时，那么dp[i][j]肯定不是回文串，所以就是false
当s[i]等于s[j]时,有以下几种情况：
1.i==j只有一个字母，那就是回文
2.i==j-1,相邻的字母，那肯定也是
3.j-i>1,相差好几个字母，此时判断dp[i+1][j-1]是不是回文就行，如果是，加上前提条件，肯定也是回文
初始化：默认为false即可
遍历顺序有很大差别，因为dp[i+1][j-1]在二维数组里面，在dp[i][j]的右下角，dp[i][j]依赖dp[i+1][j-1]的数据，
所以二维数组的遍历就需要从下到上，从左到右进行，才可以满足条件
*/
func countSubstrings(s string) int {
	length := len(s)
	dp := make([][]bool, length)
	for i := range s {
		dp[i] = make([]bool, length)
	}

	result := 0

	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					result += 1
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					result += 1
					dp[i][j] = true
				}
			}
		}
	}

	return result
}

// 最长回文字串，不删除版本
func longestPalindromeSubSeq(s string) int {
	length := len(s)
	dp := make([][]bool, length)
	for i := range s {
		dp[i] = make([]bool, length)
	}

	long := 0

	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					fmt.Println(i, j, long)
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					fmt.Println(i, j, long)

				}
			}

			if dp[i][j] && j-i+1 > long {
				long = j - i + 1
			}
		}
	}

	return long
}

/*
516. 最长回文子序列
dp[i][j]代表以i+1,j-1段最长回文子序列的长度
当s[i]==s[j]的时候，dp[i][j]=dp[i+1][j-1]+2
不等于时，就需要判断dp[i+1][j],dp[i][j-1]哪个长度最长
初始化：i和j相同的时候，dp[i][j]=1，代表s[i]这个单独的字母是回文
遍历顺序，因为求dp[i][j]时，需要dp[i+1][j-1]，dp[i+1][j]，dp[i][j-1]
所以顺序是从下往上，从左往右进行遍历，可以画个图看看
*/
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = leetcode.Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	//这里需要根据dp的定义确定返回哪个i，j
	return dp[0][len(s)-1]
}

/*
739. 每日温度
输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
*/
func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return temperatures
	}

	stack := make([]int, len(temperatures))
	result := make([]int, len(temperatures))
	result[0] = 1
	prevMax := temperatures[0]
	prevStep := 0
	for i := 1; i < len(temperatures); i++ {
		if temperatures[i] > prevMax {
			result[i] = i - prevStep + 1
			prevMax = temperatures[i]
			prevStep = i
		}
	}

	return result
}

// 优化版本
func dailyTemperaturesOP(temperatures []int) []int {
	if len(temperatures) == 0 {
		return temperatures
	}

	stack := make([]int, len(temperatures))
	result := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			//弹出栈
			top := stack[len(stack)-1]
			result[top] = i - top
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}
