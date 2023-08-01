package code_practise

import (
	"log"
	"sync"
	"time"
)

var condition = false

func ReadTask(cond *sync.Cond, who int64) {
	log.Println("ReadTask wait condition, i am ", who)
	cond.L.Lock()
	for !condition {
		cond.Wait()
	}

	log.Println("ReadTask start exec task, i am ", who)

	cond.L.Unlock()
}

func WriteTask(cond *sync.Cond) {
	log.Println("WriteTask exec task, wait 2s change condition status")
	time.Sleep(2 * time.Second)

	cond.L.Lock()
	condition = true
	cond.L.Unlock()

	//send signal wake up all g
	cond.Broadcast()
}

func LearnSyncCond() {
	cond := sync.NewCond(&sync.Mutex{})
	go ReadTask(cond, 1)
	go ReadTask(cond, 2)
	go ReadTask(cond, 3)

	WriteTask(cond)
	time.Sleep(5 * time.Second)
}
