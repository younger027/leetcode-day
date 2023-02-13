package leetcode

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	//a := []int{3, 1}
	//t.Log(Search(a, 1))

	arr := []int{1, 2, 3}
	newArr := []*int{}
	for i, _ := range arr {
		newArr = append(newArr, &arr[i])
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}
