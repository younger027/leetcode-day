package leetcode

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
)

func initListNode(data []int) *ListNode {
	head := &ListNode{
		Val:  data[0],
		Next: nil,
	}

	start := head
	for i := 1; i < len(data); i++ {
		item := ListNode{
			Val:  data[i],
			Next: nil,
		}

		head.Next = &item
		head = head.Next

	}

	return start
}

func TestRemoveListNode(t *testing.T) {
	//CountLetterFromSlice()

	node := initListNode([]int{1, 2, 3, 4, 5})

	newNode := RemoveNthFromEnd(node, 1)

	for newNode != nil {
		fmt.Println("node.value:", newNode.Val)
		newNode = newNode.Next
	}
}

type LetterFreq map[rune]int

func CountLetterFromSlice() {
	strs := []string{"abcdn", "asre", "sdasq", "sfwfwdfvergtgb", "sdfwdvdovha"}
	res := CountLetters(strs, 3)
	if res['a'] != 4 {
		log.Fatal("should 4, get %\n", res['a'])
	} else {
		fmt.Println("success")
	}
}

func CountLetters(strs []string, concurrency int) LetterFreq {
	if concurrency > len(strs) {
		concurrency = len(strs)
	}

	ch := make(chan string)
	result := make(map[rune]int, 26)

	wg := new(sync.WaitGroup)
	rmMutex := new(sync.RWMutex)

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go count(ctx, wg, ch, result, rmMutex)
	}

	for _, s := range strs {
		fmt.Println("write to ch:", s)
		ch <- s
	}

	cancel()
	wg.Wait()

	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for _, s := range strs {
	//		ch <- s
	//	}
	//	cancel()
	//
	//}()

	fmt.Println("result:", result)
	return result
}

func count(ctx context.Context, wg *sync.WaitGroup, ch chan string, result map[rune]int, rwLock *sync.RWMutex) {
	defer wg.Done()

	for {
		select {
		case str := <-ch:
			rwLock.Lock()

			for _, r := range str {
				result[r] += 1
			}
			rwLock.Unlock()

		case <-ctx.Done():
			fmt.Println("i am done ---")
			return
		}
	}
}
