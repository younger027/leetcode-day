package code_practise

import (
	"strconv"
	"sync"
	"time"
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

			for value := range dataCh {
				//log.Println(value)
				result = append(result, value)
			}

		}()
	}

	wg.Wait()

}
