package string

/*
187. 重复的DNA序列
中等
DNA序列 由一系列核苷酸组成，缩写为 'A', 'C', 'G' 和 'T'.。
例如，"ACGAATTCCG" 是一个 DNA序列 。
在研究 DNA 时，识别 DNA 中的重复序列非常有用。

给定一个表示 DNA序列 的字符串 s ，返回所有在 DNA 分子中出现不止一次的 长度为 10 的序列(子字符串)。你可以按 任意顺序 返回答案。

示例 1：
输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
输出：["AAAAACCCCC","CCCCCAAAAA"]
示例 2：

输入：s = "AAAAAAAAAAAAA"
输出：["AAAAAAAAAA"]
提示：
0 <= s.length <= 105
s[i]=='A'、'C'、'G' or 'T'
*/

func FindRepeatedDnaSequences(s string) []string {
	var result []string
	if len(s) < 10 {
		return result
	}

	need := make(map[string]int)
	left, right := 0, 0
	for right < len(s) {
		right += 1

		if right-left == 10 {
			str := s[left:right]
			need[str] += 1

			left += 1
		}
	}

	for k, v := range need {
		if v > 1 {
			result = append(result, k)
		}
	}

	return result
}
