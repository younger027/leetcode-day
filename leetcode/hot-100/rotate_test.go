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

	Rotate(matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Print(matrix[i][j], ",")
		}
		fmt.Println("-------")
	}
}
