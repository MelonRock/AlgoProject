package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestLetterNumsCache(t *testing.T) {
	alphabetCh := make(chan string, 1)
	numberCh := make(chan int, 1)

	// 打印数字
	go func() {
		for i := 0; i < 26; i++ {
			alphabetCh <- string(rune('A' + i))
			fmt.Print(<-numberCh)
		}
	}()
	// 打印字母
	go func() {
		for i := 1; i <= 26; i++ {
			fmt.Print(<-alphabetCh)
			numberCh <- i
		}
	}()
	// 等待结束
	time.Sleep(time.Second * 2)
}

func TestLetterNums(t *testing.T) {
	alphabetCh := make(chan int)
	numberCh := make(chan int)

	go func() {
		for i := 0; i < 26; i++ {
			fmt.Printf("%v", string(rune('A'+i)))
			alphabetCh <- i
			<-numberCh
		}
	}()

	go func() {
		for i := 1; i <= 26; i++ {
			<-alphabetCh
			fmt.Printf("%v", i)
			numberCh <- i
		}
	}()

	time.Sleep(time.Second * 3)
}
