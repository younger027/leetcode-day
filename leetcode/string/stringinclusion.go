package string

/*
567. 字符串的排列
中等
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。


示例 1：
输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
示例 2：
输入：s1= "ab" s2 = "eidboaoo"
输出：false

提示：
1 <= s1.length, s2.length <= 104
s1 和 s2 仅包含小写字母
*/

func CheckInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)

	for i := range s1 {
		need[s1[i]] += 1
	}

	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right += 1

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid += 1
			}
		}

		//收缩的条件
		if right-left >= len(s1) {
			if valid == len(need) {
				return true
			}

			d := s2[left]
			left += 1
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return false
}

/*
438. 找到字符串中所有字母异位词
中等
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

提示:
1 <= s.length, p.length <= 3 * 104
s 和 p 仅包含小写字母*/

func FindAnagrams(s string, p string) []int {
	result := make([]int, 0)
	need := make(map[byte]int, len(p))
	window := make(map[byte]int)
	for i := range p {
		need[p[i]] += 1
	}

	left, right := 0, 0
	valid := 0
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c] += 1
			if window[c] == need[c] {
				valid += 1
			}
		}

		if right-left >= len(p) {
			if valid == len(need) {
				result = append(result, left)
			}

			d := s[left]
			left += 1
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid -= 1
				}
				window[d] -= 1
			}
		}
	}

	return result
}

/*
3. 无重复字符的最长子串
中等
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成*/

func LengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	window := make(map[byte]int)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := s[right]
		right += 1
		window[c] += 1
		for window[c] > 1 {
			d := s[left]
			window[d] -= 1
			left += 1
		}
		if res < right-left {
			res = right - left
		}
	}

	return res
}
