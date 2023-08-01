package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main1() {
	defer ants.Release()

	runTimes := 1000

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

	//Use the pool with a function,
	//set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()

	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}

type TaskFunc func()

//go routinue pool
type Pool interface {
	AddTask(TaskFunc)
	RunTask()
}

type GPool struct {
	MaxG   int
	ChBuff int
	Ch     chan TaskFunc
	wg     *sync.WaitGroup
}

func NewGPool(size, chanBuff int) *GPool {
	return &GPool{
		MaxG: size,
		Ch:   make(chan TaskFunc, chanBuff),
		wg:   new(sync.WaitGroup),
	}
}

func (g *GPool) AddTask(f TaskFunc) {
	//get G from pool

	//run func
	g.Ch <- f
}

func (g *GPool) RunTask() {
	g.wg.Add(g.MaxG)
	for i := 0; i < g.MaxG; i++ {
		go func(i int) {
			for f := range g.Ch {
				f()
			}
			g.wg.Done()
			fmt.Printf("g id:%d is done...\n", i)
		}(i)
	}
}

func taskFunc() {
	fmt.Print("1111\n")
}

func taskFunc1() {
	fmt.Print("22222\n")
}

func taskFunc2() {
	fmt.Print("3333\n")
}

type Car struct {
	name string
}

type Audi struct {
	ac *Car
	d  int
}

func (a *Audi) PrintCarName() {
	fmt.Println("car name:", a.ac.name)
}

func main() {

	ChanControlGNums()
	// pool := NewGPool(5, 3)
	// go pool.RunTask()

	// time.Sleep(2 * time.Second)
	// pool.AddTask(taskFunc)
	// pool.AddTask(taskFunc1)
	// pool.AddTask(taskFunc2)

	// time.Sleep(2 * time.Second)

	// close(pool.Ch)
	// pool.wg.Wait()

	// time.Sleep(20 * time.Second)

	// c := &Car{
	// 	name: "younger",
	// }

	// myPool := &sync.Pool{}
	// myPool.New = func() interface{} {
	// 	return &Audi{
	// 		ac: c,
	// 	}
	// }

	// audi := myPool.Get().(*Audi)
	// audi.PrintCarName()

	// audi1 := myPool.Get().(*Audi)
	// audi1.PrintCarName()

	// fmt.Printf("%p\n", audi.ac)
	// fmt.Printf("%p\n", audi1.ac)

	// ptr1 := reflect.ValueOf(audi.ac)
	// ptr2 := reflect.ValueOf(audi1.ac)

	// fmt.Println(unsafe.Pointer(ptr1.Pointer()))
	// fmt.Println(unsafe.Pointer(ptr2.Pointer()))
}

func PrintFunc() {
	time.Sleep(10 * time.Second)
	fmt.Println("PrintFunc-----")
}
func ChanControlGNums() {
	ch := make(chan struct{}, 5)
	taskCh := make(chan func(), 100)

	go func() {
		for i := 0; i < 6; i++ {
			taskCh <- PrintFunc
		}
	}()

	for task := range taskCh {
		if task == nil {
			return
		}
		fmt.Println("read at task-----")

		go func(task func()) {
			defer func() {
				<-ch
			}()
			ch <- struct{}{}
			fmt.Println("exec task func----")
			task()
			fmt.Println("exec task func----done")
		}(task)

	}

	time.Sleep(20 * time.Second)
}

func dealwithYtdlpTime() {
	fi, err := os.Open("/root/ytmp3/logs/getyt.log")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		data := strings.Split(string(a), " ")
		durationStr := strings.TrimRight(data[len(data)-1], "]")
		duration, _ := strconv.ParseFloat(durationStr, 64)
		if duration > 3 {
			fmt.Println(duration, data[len(data)-2])
		}
	}
}
