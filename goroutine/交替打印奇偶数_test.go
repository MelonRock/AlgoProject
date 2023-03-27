package goroutine

import (
	"fmt"
	"testing"
)

/*
用协程交替打印奇偶数
*/

func printOdd(ch1 chan int, ch2 chan int) {
	for i := 1; i <= 10; i += 2 {
		<-ch1
		fmt.Print(i)
		ch2 <- 1
	}
}

func printEven(ch1 chan int, ch2 chan int) {
	for i := 2; i <= 10; i += 2 {
		<-ch2
		fmt.Print(i)
		ch1 <- 1
	}
}

func TestOddEven(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go printOdd(ch1, ch2)
	go printEven(ch1, ch2)
	ch1 <- 1
}
