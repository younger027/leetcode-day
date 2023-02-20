package leetcode

/*
35. 搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。

示例 1:
输入: nums = [1,3,5,6], target = 5
输出: 2

示例 2:
输入: nums = [1,3,5,6], target = 2
输出: 1

示例 3:
输入: nums = [1,3,5,6], target = 7
输出: 4

提示:
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 为 无重复元素 的 升序 排列数组
-104 <= target <= 104
*/

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left := BinarySearchInsert(nums, target)
	if left == len(nums) {
		return left
	}

	if nums[left] == target {
		return left
	}

	if nums[left] > target {
		return left
	} else {
		return left + 1
	}

}
func BinarySearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (right-left)/2 + left

		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	//notice here, whether you find target at array.must return left
	//find. left means index that more left number
	//not find.left means first number that bigger than target.
	//left -1 is the last index for target
	return left
}
