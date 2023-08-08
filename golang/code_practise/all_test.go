package code_practise

import (
	"testing"
	"time"
)

func TestMapAlloc(t *testing.T) {
	//mapTest()
	printAlloc()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("start test --", i)
	// 	TestMapMemUse()
	// }
	TestMapMemUse()
	printAlloc()
	TestMapMemUse()
	printAlloc()

	time.Sleep(5 * time.Second)
}
