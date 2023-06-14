package dp

import (
	"fmt"
	leetcode "leetcode/leetcode/hot-100"
	"math"
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

//416-分割等和子集
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

//1049. 最后一块石头的重量 II
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

//动态规划题目，dp[i][j]代表最多有i个0 j个1的strs数组最大子集的大小
//dp[i][j] = max(dp[i][j], dp[i-当前str的0的个数][j-当前str的1的个数]+1)
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
