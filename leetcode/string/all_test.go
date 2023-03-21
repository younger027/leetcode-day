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
	t.Log(MinWindow("a", "aa"))

	//t.Log(CheckWindow(&m, []rune("AABC")))
	//fmt.Println(m)
}
