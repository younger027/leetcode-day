package leetcode

import "fmt"

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

2-abc   3-def   4-ghi
5-jkl   6-mno    7-pqrs
8-tuv   9-wxyz

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]
提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。

*/

func LetterCombinations(digits string) []string {
	result := make([]string, 0, 19)
	length := len(digits)
	if length == 0 {
		return result
	}

	dfs(digits, "", 0, &result)

	return result
}

func dfs(digits, curStr string, i int, result *[]string) {
	mp := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

	fmt.Printf("dfs data:%p %v %v\n", result, curStr, i)
	if i >= len(digits) {
		*result = append(*result, curStr)
		return
	}

	letters := mp[digits[i]]
	for _, letter := range letters {
		dfs(digits, curStr+string(letter), i+1, result)
	}
}

func letterCombinations11(digits string) (ans []string) {
	if len(digits) == 0 {
		return
	}
	mp := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	var f func(curStr string, i int) // i: digitIndex
	f = func(curStr string, i int) {
		if i >= len(digits) {
			ans = append(ans, curStr)
			return
		} // 收集结果

		letters := mp[digits[i]]
		for _, c := range letters {
			f(curStr+string(c), i+1)
		} // 选择字母c
	}
	f("", 0)
	return
}
