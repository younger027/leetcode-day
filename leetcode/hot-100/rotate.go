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

//54,螺旋遍历矩阵

func SpiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	result := make([]int, 0)
	leftUp := 0   //上边界
	leftDown := 0 //左边界

	rightUp := n - 1   //右边界
	rightDown := m - 1 //下边界

	for len(result) < m*n {
		//上边界小于等于下边界的时候，才可以进行遍历.左--->右
		if leftUp <= rightDown {
			for i := leftUp; i <= rightUp; i++ {
				result = append(result, matrix[leftUp][i])
			}
			leftUp += 1
		}

		//左边界小于等于右边界的时候，才可以遍历。上--->下
		if leftDown <= rightUp {
			for i := leftUp; i <= rightDown; i++ {
				result = append(result, matrix[i][rightUp])
			}
			rightUp -= 1
		}

		//上边界小于等于下边界的时候，才可以进行遍历.右--->左
		if leftUp <= rightDown {
			for i := rightUp; i >= leftDown; i-- {
				result = append(result, matrix[rightDown][i])
			}
			rightDown -= 1
		}

		//左边界小于等于右边界的时候，才可以遍历。下--->上
		if leftDown <= rightUp {
			for i := rightDown; i >= leftUp; i-- {
				result = append(result, matrix[i][leftDown])
			}
			leftDown += 1
		}
	}

	return result
}

/*
59. 螺旋矩阵 II
中等
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

示例 1：
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
示例 2：

输入：n = 1
输出：[[1]]
提示：
1 <= n <= 20*/

func GenerateMatrix(n int) [][]int {
	leftUp := 0        //上边界
	leftDown := 0      //左边界
	rightUp := n - 1   //右边界
	rightDown := n - 1 //下边界
	Val := 1

	//init result matrix
	result := make([][]int, n)
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, n)
	}

	for Val <= n*n {
		if leftUp <= rightDown {
			for i := leftDown; i <= rightUp; i++ {
				result[leftUp][i] = Val
				Val += 1
			}
			leftUp += 1
		}

		if leftDown <= rightUp {
			for i := leftUp; i <= rightDown; i++ {
				result[i][rightUp] = Val
				Val += 1
			}
			rightUp -= 1
		}

		if leftUp <= rightDown {
			for i := rightUp; i >= leftDown; i-- {
				result[rightDown][i] = Val
				Val += 1
			}
			rightDown -= 1
		}

		if leftDown <= rightUp {
			for i := rightDown; i >= leftUp; i-- {
				result[i][leftDown] = Val
				Val += 1
			}
			leftDown += 1
		}

	}

	return result
}
