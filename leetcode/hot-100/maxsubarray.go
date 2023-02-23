package leetcode

import "math"

/*
53. 最大子数组和
中等
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23


提示：
1 <= nums.length <= 105
-104 <= nums[i] <= 104

进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
*/

//solution 1
//动态规划问题.dp[i] means the max sub array end with i
//dp[i] = dp[i-1]+nums[i]
//if nums[i] > 0.dp[i]=dp[i-1]+nums[i]
//if nums[i] <= 0.dp[i]=nums[i].that is import thing
//because negative number add number a must less than a

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	depth := make([]int, len(nums))
	depth[0] = nums[0]

	res := depth[0]
	for index := 1; index < len(nums); index++ {
		if depth[index-1] > 0 {
			depth[index] = depth[index-1] + nums[index]
		} else {
			//important point:
			depth[index] = nums[index]
		}

		if depth[index] > res {
			res = depth[index]
		}
	}

	return res
}

//solution 2
//分治法

func MaxSubArray_2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return findSubArray(nums, 0, len(nums)-1)
}

func findSubArray(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (right-left)/2 + left
	return GetMaxThree(findSubArray(nums, left, mid),
		findSubArray(nums, mid+1, right),
		findCrossingSubArray(nums, left, mid, right))

}

func findCrossingSubArray(nums []int, left, mid, right int) int {
	sum := 0
	leftMax := math.MinInt64
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if leftMax < sum {
			leftMax = sum
		}
	}

	sum = 0
	rightMax := math.MinInt64
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if rightMax < sum {
			rightMax = sum
		}
	}

	return leftMax + rightMax
}

func GetMaxThree(a, b, c int) int {
	return GetMax(GetMax(a, b), c)
}
