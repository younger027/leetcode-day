package leetcode

/*
生成括号
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]


提示：

1 <= n <= 8
*/
func generateParenthesis(n int) []string {
	return dp(n)[n]
}

func dp(n int) map[int][]string {
	if n == 0 {
		return map[int][]string{0: {""}}
	}

	if n == 1 {
		return map[int][]string{0: {""}, 1: {"()"}}
	}

	lastMap := dp(n - 1)
	ret := []string{}

	for i := 0; i < n; i++ {
		inner := lastMap[i]
		outter := lastMap[n-i-1]
		for _, in := range inner {
			for _, out := range outter {
				ret = append(ret, "("+in+")"+out)
			}
		}

	}
	lastMap[n] = ret

	return lastMap
}
