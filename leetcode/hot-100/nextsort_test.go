package leetcode

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	s := []int{1, 2, 3}
	NextPermutation(s)
	fmt.Println(s)
}
