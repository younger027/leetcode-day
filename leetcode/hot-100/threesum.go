package leetcode

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。


示例 2：
输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。


示例 3：
输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。


提示：

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func ThreeSum(nums []int) [][]int {
	length := len(nums)
	QuickSort(nums, 0, length-1)

	var result [][]int

	for i := 0; i < length-2; i++ {
		if nums[i] > 0 {
			return result
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L := i + 1
		R := length - 1

		for L < R {
			if nums[i]+nums[L]+nums[R] == 0 {
				var res []int
				res = append(res, nums[i], nums[L], nums[R])
				result = append(result, res)
				for L+1 < length && nums[L] == nums[L+1] {
					L++
				}

				for L < R && nums[R] == nums[R-1] {
					R--
				}

				L++
				R--
			} else if nums[i]+nums[L]+nums[R] > 0 {
				R--
			} else {
				L++
			}
		}
	}

	return result
}

func QuickSort(nums []int, start, end int) {
	if end < start {
		return
	}

	i := start
	j := end

	target := start
	for i < j {
		for i < j && nums[i] < nums[target] {
			i++
		}

		for i < j && nums[j] >= nums[target] {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[i], nums[target] = nums[target], nums[i]

	QuickSort(nums, start, i-1)
	QuickSort(nums, j+1, end)
}

/*
结题思路
1.排除特例。数组长度<3,返回空
2.排序。选择快排，平均时间复杂度是O(nlogn)
3.遍历排序后的数组：
	-如果nums[i]>0,就可以退出了，后面的数字比其小，再加肯定大于0
	-去重，避免有重复解
	-i,L=i+1,R=len-1,头，尾双指针往中间遍历，L<R
	---如果nums[i]+nums[L]+nums[R] = 0，保存结果, L++ R--。判断L，R的下一位是不是重复，如果遇到重复值，跳过。
		此时得注意边界值控制，不要超过数组左右边界。
	---如果nums[i]+nums[L]+nums[R] > 0,说明nums[R]太大了，要左移，控制边界到L
	---如果nums[i]+nums[L]+nums[R] < 0,说明nums[L]太小了，要右移，控制边界到R

*/
