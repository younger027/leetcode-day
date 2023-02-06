package leetcode

/*
32. Longest Valid Parentheses
Given a string containing just the characters '(' and ')', return the length of the longest valid (well-formed) parentheses
substring

Example 1:

Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".
Example 2:

Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".
Example 3:

Input: s = ""
Output: 0
*/

//solution-1 stack
// 时间复杂度O(n)
// 空间复杂度O(n)
func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}

	stack := []int{-1}
	max := 0

	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				max = Max(max, i-stack[len(stack)-1])

			}
		}
	}

	return max
}

/*solution-2
在此方法中，我们利用两个计数器left和right 。首先，我们从左到右遍历字符串，对于遇到的每个 ‘(’，
我们增加left计数器，对于遇到的每个 ‘)’，我们增加right计数器。每当left计数器与right计数器相等时，我们计算当前有效字符串的长度，
并且记录目前为止找到的最长子字符串。当rightright 计数器比left计数器大时，我们将left和right计数器同时变回0。

这样的做法贪心地考虑了以当前字符下标结尾的有效括号长度，每次当右括号数量多于左括号数量的时候之前的字符我们都扔掉不再考虑，
重新从下一个字符开始计算，但这样会漏掉一种情况，就是遍历的时候左括号的数量始终大于右括号的数量，即 (() ，这种时候最长有效括号是求不出来的。

解决的方法也很简单，我们只需要从右往左遍历用类似的方法计算即可，只是这个时候判断条件反了过来：
当left计数器比right计数器大时，我们将left和right计数器同时变回0,当left计数器与right计数器相等时，
我们计算当前有效字符串的长度，并且记录目前为止找到的最长子字符串,这样我们就能涵盖所有情况从而求解出答案。

时间复杂度O(n)
空间复杂度O(1)
*/
func longestValidParentheses2(s string) int {
	if len(s) <= 1 {
		return 0
	}

	left, right, max := 0, 0, 0
	for _, c := range s {
		if c == '(' {
			left += 1
		} else {
			right += 1
		}

		if right == left {
			max = Max(max, 2*left)
		}

		if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left += 1
		} else {
			right += 1
		}

		if right == left {
			max = Max(max, 2*left)
		}

		if left > right {
			left, right = 0, 0
		}
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
