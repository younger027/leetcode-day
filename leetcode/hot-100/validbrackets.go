package leetcode

//解题思路：左括号放入数据中，遇到右括号的时候，他就==数组的最后一个左括号对应的右括号。
//需要注意的点是，加入了一个？号，作为边界处理。
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	dict := map[byte]byte{'{': '}', '[': ']', '(': ')', '?': '?'}

	stack := make([]byte, 0)
	stack = append(stack, '?')

	for _, c := range s {
		if _, ok := dict[byte(c)]; ok {
			stack = append(stack, byte(c))
		} else {
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if dict[right] != byte(c) {
				return false
			}
		}
	}

	return len(stack) == 1
}
