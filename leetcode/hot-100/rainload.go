package leetcode

/*
接雨水
给定n个非负整数表示每个宽度为1的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

提示：
n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/

//solution 1 can't AC, find by rows.exec time out
//time complexity:O(H*N)
//space complexity:O(1)

func Trap(array []int) int {
	if len(array) == 0 {
		return -1
	}

	max := MaxOfArray(array)
	sum := 0
	for i := 1; i <= max; i++ {
		isStart := false
		temp := 0

		for j := 0; j < len(array); j++ {
			if isStart && array[j] < i {
				temp += 1
			}

			if array[j] >= i {
				sum += temp
				temp = 0
				isStart = true
			}
		}
	}

	return sum
}

func MaxOfArray(array []int) int {
	result := 0
	for _, num := range array {
		if num > result {
			result = num
		}
	}

	return result
}

//solution 2 find by column
//time complexity:O(n^2)
//space complexity:O(1)

func Trap2(array []int) int {
	if len(array) == 0 {
		return -1
	}

	sum := 0
	for i := 0; i < len(array); i++ {
		maxLeft := 0
		for j := 0; j < i; j++ {
			if maxLeft < array[j] {
				maxLeft = array[j]
			}
		}

		maxRight := 0
		for k := i + 1; k < len(array); k++ {
			if maxRight < array[k] {
				maxRight = array[k]
			}
		}

		secondHeight := GetMin(maxLeft, maxRight)
		if array[i] < secondHeight {
			sum += secondHeight - array[i]
		}
	}

	return sum
}

func GetMin(a, b int) int {
	if a < b {
		return a
	}

	return b

}

func GetMax(a, b int) int {
	if a < b {
		return b
	}

	return a
}

//solution 3 dynamic programming
//time complexity:O(n)
//space complexity:O(n)

func Trap3(array []int) int {
	if len(array) == 0 {
		return -1
	}

	//maxLeft[i] means max number that left of I
	//maxRight[i] means max number that right of I
	//max[i] do not contain index I number
	maxLeft := make([]int, len(array))
	maxRight := make([]int, len(array))
	sum := 0

	for i := 1; i < len(array)-1; i++ {
		maxLeft[i] = GetMax(maxLeft[i-1], array[i-1])
	}

	for i := len(array) - 2; i >= 0; i-- {
		maxRight[i] = GetMax(maxRight[i+1], array[i+1])
	}

	for i := 0; i < len(array); i++ {
		secondHeight := GetMin(maxLeft[i], maxRight[i])
		if array[i] < secondHeight {
			sum += secondHeight - array[i]
		}
	}

	return sum
}

/*
假设一开始left-1大于right+1，则之后right会一直向左移动，直到right+1大于left-1。
在这段时间内right所遍历的所有点都是左侧最高点maxleft大于右侧最高点maxright的，
所以只需要根据原则判断maxright与当前高度的关系就行。反之left右移，所经过的点只要判断maxleft与当前高度的关系就行。

time complexity:O(n)
space complexity:O(1)
*/

func Trap4(array []int) int {
	if len(array) == 0 {
		return -1
	}

	left := 1
	right := len(array) - 2
	sum := 0
	maxLeft, maxRight := 0, 0

	for i := 1; i < len(array)-1; i++ {
		if array[left-1] < array[right+1] {
			maxLeft = GetMax(maxLeft, array[left-1])
			min := maxLeft
			if min > array[left] {
				sum += min - array[left]
			}
			left += 1
		} else {
			maxRight = GetMax(maxRight, array[right+1])
			min := maxRight
			if min > array[right] {
				sum += min - array[right]
			}

			right -= 1
		}

	}

	return sum
}
