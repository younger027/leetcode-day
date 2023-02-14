package leetcode

/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]

提示：
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109

*/

//solution 1 by myself,时间复杂度O(N),空间复杂度O(1)

func SearchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}

	if len(nums) == 1 && nums[0] == target {
		result = []int{0, 0}
		return result
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			temp := mid
			for temp >= 0 && nums[temp] == target {
				result[0] = temp
				temp -= 1
			}

			temp = mid
			for temp <= len(nums)-1 && nums[temp] == target {
				result[1] = temp
				temp += 1
			}
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

//solution 2 by official,时间复杂度O(logN),空间复杂度O(1)

func SearchRange2(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}

	leftIndex := BinarySearch(nums, target)
	rightIndex := BinarySearch(nums, target+1) - 1

	if leftIndex <= rightIndex && rightIndex < len(nums) && nums[leftIndex] == target && nums[rightIndex] == target {
		return []int{leftIndex, rightIndex}
	}
	return result
}

//BinarySearch find first more than target number index
func BinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right-left)/2 + left

		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	//notice here, either find target at array.must return left
	//find. left means index that more left number
	//not find.left means first number that bigger than target.
	//left -1 is the last index for target
	return left
}
