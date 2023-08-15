package interview

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestInitTreeNode(t *testing.T) {
	// node := InitTreeNode([]int{1, 2, 3, 4, 5, 6, 7}, 0)
	// t.Log(node)

	// t.Log(BinartTreePre(node))
	// t.Log(BinartTreeMiddle(node))
	// t.Log(BinartTreeBack(node))
	// t.Log(BinartTreeLevel(node))

	// RecursiveBinartTreePre(node)
	// t.Log("\n")
	// RecursiveBinartTreeMiddle(node)
	// t.Log("\n")
	// RecursiveBinartTreeBack(node)
	// t.Log("\n")
	// RecursiveBinartTreeLevel(node)

	// t.Log(numRollsToTarget(1, 6, 3))
	// t.Log(numRollsToTarget(2, 6, 7))

	// t.Log(mergeAlternately("abc", "pqr"))

	// t.Log(0%4, 1%4, 2%4, 3%4, 4%4)
	// t.Log(gcdOfStringsMain("ABABAB", "ABAB"))
	// t.Log(gcdOfStringsMain("ABCABC", "ABC"))
	// t.Log(gcdOfStringsMain("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"))
	// t.Log(gcdOfStringsMain("AA", "A"))
	// t.Log(gcdOfStringsMain("AAAAAAAAA", "AAACCC"))
	//t.Log(gcdOfStrings("ABCDEF", "ABC"))
	//t.Log(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	//t.Log(canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
	//t.Log(reverseVowels("hello"))
	//t.Log(reverseWords("a good   example"), "-")
	//t.Log(reverseWords("a good   example"), "-")
	//t.Log(productExceptSelfOp([]int{-1, 1, 0, -3, 3}))

	//t.Log(increasingTriplet([]int{1, 2, 3, 4, 5}))

	slice := make([]int, 0, 3)
	slice = append(slice, 1, 2, 3)
	fmt.Println(unsafe.Pointer(&slice[0]))
	testSlice(slice)
	fmt.Println(slice)
}

func SliceAppend(data []int) {
	data = append(data, 1, 2)
	fmt.Println("slice append ", data)
}

func testSlice(slice []int) {
	fmt.Println(unsafe.Pointer(&slice[0]))
	slice = append(slice, 4)
	slice[0] = 10
	fmt.Println(unsafe.Pointer(&slice[0]))
	fmt.Println("----", slice[3])
	fmt.Println(slice)
}

func TestCompress(t *testing.T) {
	// t.Log(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	// t.Log(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c', 'd'}))
	// t.Log(compress([]byte{'a'}))
	// t.Log(compress([]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}))

	//moveZeroes([]int{0, 1, 0, 3, 12})
	//t.Log(isSubsequence("axc", "ahbgdc"))

	//t.Log(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	// t.Log(maxOperations([]int{2, 2, 2, 3, 1, 1, 4, 1}, 4))
	// t.Log(maxOperations([]int{3, 1, 3, 4, 3}, 6))
	// t.Log(maxOperations([]int{1, 2, 3, 4}, 5))

	t.Log(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
