package leetcode

/*
请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）

注意：
一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
空白格用 '.' 表示。

示例 1：
输入：board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
输出：true

示例 2：
输入：board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
输出：false
解释：除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。 但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。


提示：

board.length == 9
board[i].length == 9
board[i][j] 是一位数字（1-9）或者 '.'

*/

//isValidSudoku solving ideas
// judge row whether exist same num
// judge column whether exist same num
// judge 3x3 sudoku whether meet the conditions

func IsValidSudoku(board [][]byte) bool {
	//row column box first array index means it belong where row column and box
	//second index array means the index of first array(row column box) whether existed before
	row := make([][]int, 9)
	column := make([][]int, 9)
	box := make([][]int, 9)

	for i := 0; i < 9; i++ {
		row[i] = make([]int, 9)
		column[i] = make([]int, 9)
		box[i] = make([]int, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			nowNum := board[i][j] - '1'
			if row[i][nowNum] == 1 || column[j][nowNum] == 1 {
				return false
			}

			if box[j/3+(i/3)*3][nowNum] == 1 {
				return false
			}
			row[i][nowNum] = 1
			column[j][nowNum] = 1
			box[j/3+(i/3)*3][nowNum] = 1
		}
	}

	return true
}