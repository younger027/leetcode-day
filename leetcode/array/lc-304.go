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
	return this.preNum[row2][col2] - this.preNum[row2][col1-1] - this.preNum[row1-1][col2] + this.preNum[row1-1][col1]
}

func (this *NumMatrix) Show() {
	for i := 0; i < len(this.preNum); i++ {
		for j := 0; j < len(this.preNum[0]); j++ {
			//calculate num for 0,0--i,j
			fmt.Print(this.preNum[i][j], ",")
		}
		fmt.Print("\n")
	}
}
