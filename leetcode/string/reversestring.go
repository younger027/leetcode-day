package string

import "fmt"

//hello i am younger--->younger am i hello
/*

输入：s = "the sky is blue"
输出："blue is sky the"
示例 2：

输入：s = "  hello world  "
输出："world hello"
解释：反转后的字符串中不能存在前导空格和尾随空格。
示例 3：

输入：s = "a good   example"
输出："example good a"
*/

func ReverseString(s string) string {
	if len(s) <= 1 {
		return s
	}

	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	fmt.Print("ssss:", string(r))
	last := 0
	for k := 0; k < len(r); k++ {
		if r[k] == ' ' || k == len(r)-1 {
			if k == len(r)-1 {
				k += 1
			}
			for i, j := last, k-1; i < (k-last)/2+last; i, j = i+1, j-1 {
				r[i], r[j] = r[j], r[i]
			}
			last = k + 1
		}
	}

	return string(r)
}
