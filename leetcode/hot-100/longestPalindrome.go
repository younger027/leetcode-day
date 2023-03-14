package leetcode

/*
5.Longest Palindromic Substring

Given a string s, return the longest palindromic substring in s.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"


Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.
*/

func LongestPalindrome(s string) string {
	length := len(s)

	if length == 1 {
		return string(s[0])
	}

	if length == 2 && s[0] != s[1] {
		return string(s[0])
	}

	start := 0
	end := 0
	for i := 0; i < length-1; i++ {
		j := length - 1
		for i < j {
			if s[i] == s[j] {
				isPalindrome := findPalindrome(s, i, j)
				if isPalindrome && (j-i) > (end-start) {
					start = i
					end = j
					continue
				}
			}

			j -= 1
		}

	}

	return s[start : end+1]
}

func findPalindrome(s string, i, j int) bool {
	if j <= i {
		return false
	}

	for i <= j {
		if s[i] == s[j] {
			i += 1
			j -= 1
			continue
		}

		return false
	}

	return true
}

//solution2.从中间向两边扩展。如果长度是奇数的时候他只有一个中间字符，如果是偶数的时候，可以认为他有2个中心字符。
func LongestPalindrome2(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		s1 := Palindrome(s, i, i)
		s2 := Palindrome(s, i, i+1)

		if len(res) < len(s1) {
			res = s1
		}

		if len(res) < len(s2) {
			res = s2
		}
	}

	return res
}

func Palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l -= 1
		r += 1
	}

	return s[l+1 : r]
}
