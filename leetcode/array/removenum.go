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
