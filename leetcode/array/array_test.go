package array

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	array := []int{2, 0, 1, 0, 3, 12}
	//t.Log(RemoveElement(array, 1))
	//t.Log(RemoveElementByOrder(array, 1))

	//t.Log(RemoveDuplicates(array))
	MoveZeroes(array)

	for i := 0; i < len(array); i++ {
		fmt.Println("node value---", array[i])
	}
}

func TestBackspaceCompare(t *testing.T) {
	//BackspaceCompare("ab##", "")
	t.Log(BackspaceCompareO1("ab#c", "ad#c"))
}

func TestConstructorTwo(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5}}
	this := ConstructorTwo(matrix)
	this.Show()

	t.Log(this.SumRegion(2, 1, 4, 3))
	t.Log(this.SumRegion(1, 1, 2, 2))
	t.Log(this.SumRegion(1, 2, 2, 4))
}
