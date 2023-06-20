package main

import (
	"fmt"
	"github.com/arl/statsviz"
	"github.com/shirou/gopsutil/process"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

func TestCloseChan() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the moderator goroutine shown below.
	// Its reveivers are all senders and receivers of dataCh.
	toStop := make(chan string, 1)
	// the channel toStop is used to notify the moderator
	// to close the additional signal channel (stopCh).
	// Its senders are any senders and receivers of dataCh.
	// Its reveiver is the moderator goroutine shown below.

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop // part of the trick used to notify the moderator
		// to close the additional signal channel.
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				Val := rand.Intn(MaxRandomNumber)
				if Val == 0 {
					// here, a trick is used to notify the moderator
					// to close the additional signal channel.
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// the first select here is to try to exit the
				// goroutine as early as possible.
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- Val:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				select {
				case Val := <-dataCh:
					//do something
					log.Println(Val)
				default:
				}

				// same as senders, the first select here is to
				// try to exit the goroutine as early as possible.
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case Val := <-dataCh:
					if Val == MaxRandomNumber-1 {
						// the same trick is used to notify the moderator
						// to close the additional signal channel.
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(Val)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}

func printGcStats() {
	t := time.NewTicker(time.Second)
	s := debug.GCStats{}

	for {
		select {
		case <-t.C:
			fmt.Printf("ticker coming--------")
			debug.ReadGCStats(&s)
			fmt.Printf("gc %d last@%v, PauseTotal %v\n", s.NumGC, s.LastGC, s.PauseTotal)
		}
	}
}

func main() {
	// 实时查看 Go 应用程序运行时统计信息(GC，MemStats 等)
	statsviz.RegisterDefault()

	go printGcStats()

	go func() {
		fmt.Println(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", "6060"), nil))
	}()

	//打印cpu和内存的使用信息
	go func() {
		pid := os.Getpid()
		fmt.Println("当前程序的进程号为：", pid)

		p, _ := process.NewProcess(int32(pid))
		for {
			v, _ := p.CPUPercent()
			if v > 0 {
				memPercent, _ := p.MemoryPercent()
				fmt.Printf("该进程的cpu占用率:%v,内存占用:%v, 时间:%v\n", v, memPercent, time.Now().Format("2006-01-02 15:04:05"))
				println("---------------分割线------------------")
			}
			time.Sleep(5 * time.Second)
		}

	}()

	fmt.Printf("最初！程序中的goroutine数量为：%d\n", runtime.NumGoroutine())
	for i := 0; i < 1000000; i++ {
		go func() {
			time.Sleep(time.Second * 10)
		}()
	}

	fmt.Printf("for循环结束后！程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	time.Sleep(5e9)

	fmt.Printf("5s后程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	time.Sleep(5e9)

	fmt.Printf("10s后程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	time.Sleep(5e9)
	fmt.Printf("15s后程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	time.Sleep(5e9)
	fmt.Printf("20s后程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	// 用于阻塞不使程序退出
	select {}
}
