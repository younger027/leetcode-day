package leetcode

/*
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

提示：
n == matrix.length == matrix[i].length
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000

*/
func Rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			//be careful rotate left or right
			//temp := matrix[i][j]
			//matrix[i][j] = matrix[n-1-j][i]
			//matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			//matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			//matrix[j][n-1-i] = temp

			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]

		}
	}
}

//顺时针90
func RotateMatrix(matrix [][]int) {
	m := len(matrix)

	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	i := 0
	j := m - 1

	for i < j {
		for k := 0; k < m; k++ {
			matrix[k][i], matrix[k][j] = matrix[k][j], matrix[k][i]
		}
		i += 1
		j -= 1
	}
}

//逆时针
func RotateMatrix2(matrix [][]int) {
	m := len(matrix)

	for i := 0; i < m; i++ {
		for j := i; j < m-1-i; j++ {
			matrix[i][j], matrix[m-1-j][m-1-i] = matrix[m-1-j][m-1-i], matrix[i][j]
		}
	}

	i := 0
	j := m - 1

	for i < j {
		for k := 0; k < m; k++ {
			matrix[k][i], matrix[k][j] = matrix[k][j], matrix[k][i]
		}
		i += 1
		j -= 1
	}
}

func SpiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	upper_bound, lower_bound := 0, m-1
	left_bound, right_bound := 0, n-1
	res := make([]int, 0, m*n)
	// len(res) == m * n 则遍历完整个数组
	for len(res) < m*n {
		if upper_bound <= lower_bound {
			// 在顶部从左向右遍历
			for j := left_bound; j <= right_bound; j++ {
				res = append(res, matrix[upper_bound][j])
			}
			// 上边界下移
			upper_bound++
		}

		if left_bound <= right_bound {
			// 在右侧从上向下遍历
			for i := upper_bound; i <= lower_bound; i++ {
				res = append(res, matrix[i][right_bound])
			}
			// 右边界左移
			right_bound--
		}

		if upper_bound <= lower_bound {
			// 在底部从右向左遍历
			for j := right_bound; j >= left_bound; j-- {
				res = append(res, matrix[lower_bound][j])
			}
			// 下边界上移
			lower_bound--
		}

		if left_bound <= right_bound {
			// 在左侧从下向上遍历
			for i := lower_bound; i >= upper_bound; i-- {
				res = append(res, matrix[i][left_bound])
			}
			// 左边界右移
			left_bound++
		}
	}
	return res
}
