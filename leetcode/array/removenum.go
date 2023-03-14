package array

func RemoveElement(nums []int, val int) int {

	/*
		left := 0
		    for _, v := range nums { // v 即 nums[right]
		        if v != val {
		            nums[left] = v
		            left++
		        }
		    }
		    return left

	*/
	j := len(nums) - 1
	for i := 0; i < j; i++ {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			i -= 1
			j -= 1
		}
	}

	return j + 1
}

func RemoveElementByOrder(nums []int, val int) int {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[slow] = nums[i]
			slow += 1
		}
	}

	return slow
}

/*

26. 删除有序数组中的重复项
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。

由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。

将最终结果插入 nums 的前 k 个位置后返回 k 。

不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
*/

func RemoveDuplicates(nums []int) int {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[slow] {
			slow += 1
			nums[slow] = nums[i]
		}
	}

	return slow + 1
}

/*
283. 移动零
简单
相关企业
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。*/

func MoveZeroes(nums []int) {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[slow], nums[i] = nums[i], nums[slow]
			slow += 1
		}

	}

}

/*
844. 比较含退格的字符串
简单

给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。
注意：如果对空文本输入退格字符，文本继续为空。

示例 1：
输入：s = "ab#c", t = "ad#c"
输出：true
解释：s 和 t 都会变成 "ac"。
示例 2：

输入：s = "ab##", t = "c#d#"
输出：true
解释：s 和 t 都会变成 ""。
示例 3：

输入：s = "a#c", t = "b"
输出：false
解释：s 会变成 "c"，但 t 仍然是 "b"。

提示：

1 <= s.length, t.length <= 200
s 和 t 只含有小写字母以及字符 '#'

进阶：

你可以用 O(n) 的时间复杂度和 O(1) 的空间复杂度解决该问题吗？
*/

//solution 1
//time complex O(m+n)
//spave complex O(m+n)

func BackspaceCompare(s string, t string) bool {

	return getString(s) == getString(t)
}

func getString(s string) string {
	bz := []rune{}

	for _, c := range s {
		if c != '#' {
			bz = append(bz, c) // 模拟入栈
		} else if len(bz) > 0 { // 栈非空才能弹栈
			bz = bz[:len(bz)-1] // 模拟弹栈
		}
	}

	return string(bz)
}

//输入：s = "ab##", t = "c#d#"
//solution 1
//time complex O(m+n)
//spave complex O(1)

func BackspaceCompareO1(s string, t string) bool {
	lens := len(s) - 1
	lent := len(t) - 1

	i := lens
	j := lent

	skips := 0
	skipt := 0

	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skips += 1
				i -= 1
			} else if skips > 0 {
				i -= 1
				skips -= 1
			} else {
				break
			}
		}

		for j >= 0 {
			if t[j] == '#' {
				skipt += 1
				j -= 1
			} else if skipt > 0 {
				j -= 1
				skipt -= 1
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}

		i--
		j--

	}
	return true
}

//303-区域和检索 - 数组不可变
//preNums[i] 记录 nums[0..i-1] 的累加和

type NumArray struct {
	preNums []int
}

func Constructor(nums []int) NumArray {
	preNums := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		preNums[i] = nums[i-1] + preNums[i-1]
	}

	return NumArray{
		preNums: preNums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preNums[right+1] - this.preNums[left]
}
