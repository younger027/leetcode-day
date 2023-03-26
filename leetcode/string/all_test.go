package string

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	t.Log(ReverseString("   hello    i    am   younger    "))
}

func TestMinWindow(t *testing.T) {
	//m := map[rune]int{
	//	'A': 2,
	//	'B': 1,
	//	'C': 1,
	//}
	t.Log(MinWindow("ADOBECODEBANC", "ABC"))
	t.Log(MinWindow("a", "a"))
	t.Log(MinWindow("a", "aa"))

	//t.Log(CheckWindow(&m, []rune("AABC")))
	//fmt.Println(m)
}

func TestCheckInclusion(t *testing.T) {

	t.Log(CheckInclusion("ab", "eidboaoo"))

}

func TestFindAnagrams(t *testing.T) {

	t.Log(FindAnagrams("abab", "ab"))

}
func TestLengthOfLongestSubstring(t *testing.T) {

	t.Log(LengthOfLongestSubstring("bbbbb"))

}
