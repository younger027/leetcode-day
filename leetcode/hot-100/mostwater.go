package leetcode

import "fmt"

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。


*/

func minInter(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

func MaxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	maxWater := 0

	i, j := 0, len(height)-1
	for i < j {
		fmt.Println(i, j)
		water := minInter(height[i], height[j]) * (j - i)
		if water > maxWater {
			maxWater = water
		}

		if height[i] < height[j] {
			i++
			fmt.Println("i++++")

		} else if height[i] >= height[j] {
			fmt.Println("j------")
			j--
		}
	}

	return maxWater
}
