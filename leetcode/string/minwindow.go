package string

import (
	"fmt"
	"math"
)

/*
76. 最小覆盖子串
困难
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
示例 2：

输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。

提示：
m == s.length
n == t.length
1 <= m, n <= 105
s 和 t 由英文字母组成

进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
*/

//自己写的算法，可以通过case。但是时间超时。
func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	start := 0
	length := math.MaxInt32

	window := make(map[rune]int, 26)
	ss := []rune(s)
	tt := []rune(t)

	left, right := 0, 0
	for right < len(ss) {
		c := s[right]
		window[rune(c)] += 1
		right += 1

		//检查window是否满足需求
		if !CheckWindow(&window, tt) {
			continue
		}

		fmt.Printf("window: [%d, %d)\n", left, right)
		fmt.Printf("result: [%d, %d)\n", start, length)
		for right > left {
			if length > right-left {
				start = left
				length = right - start
				fmt.Printf("result-----: [%d, %d)\n", start, length)
			}

			leftCh := rune(s[left])
			if window[leftCh] > 0 {
				window[leftCh] -= 1
			}
			left += 1

			if !CheckWindow(&window, tt) {
				break
			}

		}

	}

	if length == math.MaxInt32 {
		return ""
	}

	return s[start : start+length]
}

func CheckWindow(window *map[rune]int, t []rune) bool {
	tMap := make(map[rune]int, len(t))
	for i := 0; i < len(t); i++ {
		tMap[t[i]] += 1

	}

	for k, v := range tMap {
		count, ok := (*window)[k]
		if ok && v <= count {
			continue
		} else {
			return false
		}
	}

	return true
}

//代码随想录的算法参考

func MinWindowTwo(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	start := 0
	length := math.MaxInt32
	window := make(map[byte]int) //不必存所有的字符出现的次数。只存t存在的即可
	need := make(map[byte]int, len(t))
	for i := range t {
		need[t[i]] += 1
	}

	valid := 0 // 判断窗口中是否已经包含了字串 t 中所有字符
	left, right := 0, 0

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}

			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	}

	return s[start : start+length]
}
