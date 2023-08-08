package code_practise

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

func SafeCloseChanOneSender() {
	numReceiver := 10

	wgReceiver := &sync.WaitGroup{}

	closed := make(chan struct{})
	closing := make(chan struct{})
	dataCh := make(chan int, 10)

	go func() {
		time.Sleep(3 * time.Second)

		select {
		case closing <- struct{}{}:
			//wait sender goroutine exit.
			<-closed
		case <-closed:
		}

	}()

	//sender
	go func() {
		defer func() {
			close(closed)
			close(dataCh)
		}()

		i := 0
		for {
			//quick closing
			select {
			case <-closing:
				return
			default:
			}

			select {
			case <-closing:
				return
			case dataCh <- i:
				i += 1
			}
		}
	}()

	wgReceiver.Add(numReceiver)

	for i := 0; i < numReceiver; i++ {
		go func() {
			defer wgReceiver.Done()
			for {
				select {
				case <-closed:
					return
				case <-dataCh:
					//fmt.Println("read data---", d)
				}
			}

		}()
	}

	wgReceiver.Wait()

}

func SafeCloseChanMSender() {
	numReceivers := 10
	numSenders := 10

	wg := &sync.WaitGroup{}
	wg.Add(numSenders + numReceivers)

	closed := make(chan struct{})
	closing := make(chan struct{})
	dataCh := make(chan int)
	middleCh := make(chan int) //never close.wait all goroutine exit.GC deal with middleCh

	stopCode := func() {
		select {
		case closing <- struct{}{}:
			//wait sender goroutine exit.
			<-closed
		case <-closed:
		}

	}

	for i := 0; i < 10; i++ {
		go func(id string) {
			r := 3
			time.Sleep(time.Duration(r) * time.Second)
			stopCode()
		}(strconv.Itoa(i))
	}

	go func() {
		exitCode := func(v int, needSend bool) {
			close(closed)
			if needSend {
				dataCh <- v
			}
			close(dataCh)
		}

		for {
			//quick closing
			select {
			case <-closing:
				exitCode(1, false)
				return
			case data := <-middleCh:
				select {
				case <-closing:
					exitCode(1, true)
					return
				case dataCh <- data:
				}

			}
		}
	}()

	//sender
	for i := 0; i < numSenders; i++ {
		go func() {
			defer wg.Done()

			data := 0
			for {
				//quick closing
				select {
				case <-closed:
					return
				default:
				}

				select {
				case <-closed:
					return
				case middleCh <- data:
					data += 1
				}
			}
		}()

	}

	result := make([]int, 10)
	for i := 0; i < numReceivers; i++ {
		go func() {
			defer wg.Done()

			for Val := range dataCh {
				//log.Println(Val)
				result = append(result, Val)
			}

		}()
	}

	wg.Wait()

}

const N = 128

func randBytes() [N]byte {
	return [N]byte{}
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

func mapTest() {
	n := 1_000_000
	m := make(map[int][N]byte, 0)
	printAlloc()

	var B uint8
	for i := 0; i < n; i++ {
		curB := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int)(unsafe.Pointer(&m)))) + 9))
		if B != curB {
			fmt.Println(curB)
			B = curB
		}

		m[i] = randBytes()
	}
	printAlloc()

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
}

var mm map[int][N]byte

func TestMapMemUse() {
	n := 1_000_000
	mm = make(map[int][N]byte)

	for i := 0; i < n; i++ {
		mm[i] = randBytes()
	}
	runtime.GC()
}
