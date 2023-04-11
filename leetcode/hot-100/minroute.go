package leetcode

//64最小路径和
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = minNum(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			}
		}
	}

	return dp[m-1][n-1]
}

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

/*72 编辑距离
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

示例 1：
输入：word1 = "horse", word2 = roe "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

提示：
0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成


解题思路：
状态转移方程
定义dp[i][j],代表从word1的i位置，满足word2的j位置所用的最小的改动次数

那么当word1[i] == word2[j]时，dp[i][j] = dp[i-1][j-1]
当条件不成立时，word1[i] != word2[j]。我们思考下适用替换，删除，增加的情况是什么样子。
替换类似：roe "ros"
当word1的i-1之前和word2的j-1之前都一样，只要替换掉word1的i位置就行。所以此时是dp[i-1][j-1]+1

增加：
当word1的i位置之前和word2的j-1之前都一样，只要将word2[j]加到word1后面就行。dp[i][j-1]+1

删除：
当word1的i-1位置和word2的j之前都一样，那么只要将word1的i位置删除即可。dp[i-1][j]+1

特殊情况
对word1为空，word2不为空的情况，只能增加
对word1不为空，word2为空的情况，只能删除
*/

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	dp := make([][]int, l1+1)

	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}

	//列
	for i := 1; i <= l1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	//行
	for j := 1; j <= l2; j++ {
		dp[0][j] = dp[0][j-1] + 1
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
