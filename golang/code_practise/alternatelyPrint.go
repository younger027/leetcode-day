package code_practise

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
//
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

func printNumber(wg *sync.WaitGroup, chInt chan int, chLetter chan int) {
	num := 1

	defer wg.Done()
	defer close(chLetter)

	for {
		select {
		case <-chInt:
			fmt.Printf("%d%d", num, num+1)
			num = num + 2

			if num > 28 {
				fmt.Println("close num-----")
				return
			}

			chLetter <- 1
		}
	}

}

func printLetters(wg *sync.WaitGroup, chInt chan int, chLetter chan int) {
	defer wg.Done()
	defer close(chInt)

	letter := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	i := 0
	for {
		select {
		case <-chLetter:
			fmt.Print(letter[i : i+2])

			if i >= len(letter)-2 {
				fmt.Println("iiiiiii-", i)
				return
			}

			i = i + 2
			chInt <- 1

		}
	}

}

func AlternatelyPrintNumAndLetter() {
	chInt := make(chan int)
	chLetter := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go printNumber(wg, chInt, chLetter)
	go printLetters(wg, chInt, chLetter)

	chInt <- 1

	wg.Wait()

}

func WriteChan(wg *sync.WaitGroup, ctx context.Context, ch chan int, i int, no int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteChan ctx done, routine exit, no.", no)
			return
		default:
			ch <- i
			fmt.Println("write data:", i)
			i++
		}
	}
}

func ReadChan(wg *sync.WaitGroup, ctx context.Context, ch chan int, no int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("ReadChan ctx done, routine exit, no.", no)
			return
		case i := <-ch:
			fmt.Println("read data:", i)
		}
	}
}

func MutifyWriteReadChanOP() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	wg.Add(6)
	for i := 0; i < 3; i++ {
		go WriteChan(wg, ctx, ch, i, i)
	}

	for i := 0; i < 3; i++ {
		go ReadChan(wg, ctx, ch, i)
	}

	wg.Wait()
	close(ch)
}
