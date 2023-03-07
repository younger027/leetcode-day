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
