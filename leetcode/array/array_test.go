package array

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	array := []int{1}
	t.Log(RemoveElement(array, 1))

	for i := 0; i < len(array); i++ {
		fmt.Println("node value---", array[i])
	}
}
