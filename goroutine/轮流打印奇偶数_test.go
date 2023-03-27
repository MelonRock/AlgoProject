package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

/*
首先因为变量ch是一个无缓冲的channel, 所以只有读写同时就绪时才不会阻塞。
所以两个goroutine会同时进入各自的 if 语句（此时 i 是相同的），
但是此时只能有一个 if 是成立的，不管哪个goroutine快，
都会由于读channel或写channel导致阻塞
*/

func printOdd1(flag chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		flag <- 1
		if i%2 == 1 {
			fmt.Print(i)
		}
	}
}

func printEven1(flag chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		<-flag
		if i%2 == 0 {
			fmt.Print(i)
		}
	}
}

func TestOddAdd1(t *testing.T) {
	var flag = make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go printOdd1(flag, &wg)
	go printEven1(flag, &wg)
	wg.Wait()
}
