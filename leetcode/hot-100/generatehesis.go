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

题目解析：
归类：动态规划。学习参考链接：https://www.zhihu.com/question/23995189
f(n-1)代表n-1对括号的最全组合。那f(n)="(" + f(i) + ")" + f(n-i-1)

当我们清楚所有 i<n 时括号的可能生成排列后，对与 i=n 的情况，我们考虑整个括号排列中最左边的括号。
它一定是一个左括号，那么它可以和它对应的右括号组成一组完整的括号 "( )"，我们认为这一组是相比 n-1 增加进来的括号。

那么，剩下 n-1 组括号有可能在哪呢？

【这里是重点，请着重理解】
剩下的括号要么在这一组新增的括号内部，要么在这一组新增括号的外部（右侧）。

既然知道了 i<n 的情况，那我们就可以对所有情况进行遍历：

"(" + 【i=p时所有括号的排列组合】 + ")" + 【i=q时所有括号的排列组合】

其中 p + q = n-1，且 p q 均为非负整数。

事实上，当上述 p 从 0 取到 n-1，q 从 n-1 取到 0 后，所有情况就遍历完了。

注：上述遍历是没有重复情况出现的，即当 (p1,q1)≠(p2,q2) 时，按上述方式取的括号组合一定不同。
即当 (p1,q1)==(p2,q2)时，结果会有重复。但是在map中，n对应的结果都是一样的。只是多计算了几次而已
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
