package leetcode

import (
	"fmt"
	"math"
)

/*
4. 寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

示例 2：
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-10^6 <= nums1[i], nums2[i] <= 10^6
*/

//解法一
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	min1 := (m + n + 1) / 2
	min2 := (m + n + 2) / 2

	a := findKth(nums1, 0, nums2, 0, min1)
	b := findKth(nums1, 0, nums2, 0, min2)
	fmt.Println(a, b)
	return (float64(a) + float64(b)) / 2.0
}

func findKth(nums1 []int, i int, nums2 []int, j, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}

	if j >= len(nums2) {
		return nums1[i+k-1]
	}

	if k == 1 {
		return minNum(nums1[i], nums2[j])
	}

	min1 := math.MaxInt64
	min2 := math.MaxInt64
	if i+k/2-1 < len(nums1) {
		min1 = nums1[i+k/2-1]
	}

	if j+k/2-1 < len(nums2) {
		min2 = nums2[j+k/2-1]
	}

	if min1 < min2 {
		return findKth(nums1, i+k/2, nums2, j, k-k/2)
	}

	return findKth(nums1, i, nums2, j+k/2, k-k/2)
}

func minNum(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func minNumThree(a, b, c int) int {
	if a > b {
		if c > b {
			return b
		} else {
			return c
		}
	} else {
		if c > a {
			return a
		} else {
			return c
		}
	}
}

func MaxNumThree(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if c > b {
			return c
		} else {
			return b
		}
	}
}

//解法二
func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	min1 := (m + n + 1) / 2
	min2 := (m + n + 2) / 2

	a := findKth2(nums1, 0, m, nums2, 0, n, min1)
	b := findKth2(nums1, 0, m, nums2, 0, n, min2)

	return (float64(a) + float64(b)) * 0.5
}

func findKth2(nums1 []int, s1, e1 int, nums2 []int, s2, e2, k int) int {
	if k/2-1 > (e1 - s1) {
		return nums2[k-(e1-s1)-1]
	}

	if k/2-1 > (e2 - s2) {
		return nums1[k-(e2-s2)-1]
	}

	if k == 1 {
		return minNum(nums1[s1], nums2[s2])
	}

	if nums1[k/2-1] > nums2[k/2-1] {
		return findKth2(nums1, s1, e1, nums2, k/2-1, e2, k-(e2-s2)/2)
	}

	return findKth2(nums1, k/2-1, e1, nums2, s2, e2, k-(e1-s1)/2)
}

/*
128. 最长连续序列
中等
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
*/

func longestConsecutive(nums []int) int {
	maxMap := make(map[int]bool)
	max := 0

	for i := 0; i < len(nums); i++ {
		maxMap[nums[i]] = true
	}

	//dp[i] 代表以i数字开始的最长连续序列长度
	//dp[i]= maxMap[i-1] >0 ?maxMap+1, 1
	for _, number := range nums {
		if !maxMap[number-1] {
			curNum := number
			curMax := 1
			for maxMap[curNum+1] {
				curMax++
				curNum++
			}

			if curMax > max {
				max = curMax
			}
		}
	}

	return max
}
