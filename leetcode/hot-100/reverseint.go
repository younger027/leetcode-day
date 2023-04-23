package leetcode

import (
	"math"
	"strconv"
	"strings"
)

//Given a signed 32-bit integer x, return x with its digits reversed.
//If reversing x causes the Val to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
//
//Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
//
//
//
//Example 1:
//
//Input: x = 123
//Output: 321
//Example 2:
//
//Input: x = -123
//Output: -321
//Example 3:
//
//Input: x = 120
//Output: 21
//
//
//Constraints:
//
//-231 <= x <= 231 - 1

//special case
// -123
// 1200
// -9 --- 9  no change
// 1534236469 after reverse more than max int

func Reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	if x >= 0 && x <= 9 {
		return x
	}

	isNegativeNum := false
	if x < 0 {
		isNegativeNum = true
		x = -x
	}

	var numSlice []string
	isMeetNoZero := false
	for {
		a := x % 10
		x = x / 10
		if a == 0 && !isMeetNoZero {
			continue
		}

		isMeetNoZero = true
		numSlice = append(numSlice, strconv.Itoa(a))

		if x == 0 {
			break
		}
	}

	result, _ := strconv.Atoi(strings.Join(numSlice, ""))
	if isNegativeNum {
		result = -result
	}

	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	return result
}
