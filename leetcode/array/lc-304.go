package array

import "fmt"

/*

304. 二维区域和检索 - 矩阵不可变
中等
477
相关企业
给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。


示例 1：



输入:
["NumMatrix","sumRegion","sumRegion","sumRegion"]
[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
输出:
[null, 8, 11, 12]

解释:
NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]);
numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)
*/

type NumMatrix struct {
	preNum [][]int
}

func ConstructorTwo(matrix [][]int) NumMatrix {
	preNums := make([][]int, len(matrix))
	for i := 0; i < len(matrix[0]); i++ {
		preNums[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			//calculate num for 0,0--i,j
			preNums[i][j] = calculate(matrix, i, j)
		}
	}

	return NumMatrix{preNum: preNums}
}

func calculate(matrix [][]int, row, col int) int {
	res := 0
	for i := 0; i <= row; i++ {
		for j := 0; j <= col; j++ {
			res += matrix[i][j]
		}
	}

	return res
}
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	fmt.Print(this.preNum[row2][col2], this.preNum[row2][col1], this.preNum[row1][col2], this.preNum[row1][col1])
	return this.preNum[row2][col2] - this.preNum[row2][col1] - this.preNum[row1][col2] + this.preNum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

//type NumMatrix struct {
//	// 定义：preSum[i][j] 记录 matrix 中子矩阵 [0, 0, i-1, j-1] 的元素和
//	preSum [][]int
//}
//
func Constructor1(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return NumMatrix{}
	}
	// 构造前缀和矩阵
	preSum := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		preSum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 计算每个矩阵 [0, 0, i, j] 的元素和
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{preNum: preSum}
}

//
//// 计算子矩阵 [x1, y1, x2, y2] 的元素和
//func (this *NumMatrix) SumRegion(x1 int, y1 int, x2 int, y2 int) int {
//	// 目标矩阵之和由四个相邻矩阵运算获得
//	fmt.Print(this.preSum[x2+1][y2+1], this.preSum[x1][y2+1], this.preSum[x2+1][y1], this.preSum[x1][y1])
//
//	return this.preSum[x2+1][y2+1] - this.preSum[x1][y2+1] - this.preSum[x2+1][y1] + this.preSum[x1][y1]
//}
//
func (this *NumMatrix) Show() {
	for i := 0; i < len(this.preNum); i++ {
		for j := 0; j < len(this.preNum[0]); j++ {
			//calculate num for 0,0--i,j
			fmt.Print(this.preNum[i][j], ",")
		}
		fmt.Print("\n")
	}
}
