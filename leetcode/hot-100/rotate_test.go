package leetcode

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	//Rotate(matrix)
	RotateMatrix2(matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Print(matrix[i][j], ",")
		}
		fmt.Println("-------")
	}
}

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	//Rotate(matrix)
	t.Log(SpiralOrder(matrix))

}

func TestGenerateMatrix(t *testing.T) {
	s := GenerateMatrix(4)
	t.Log(s)
}
