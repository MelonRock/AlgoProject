package main

import (
	"fmt"
	"testing"
)

func Producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func Consumer(id int, ch chan int, done chan bool) {
	for {
		value, ok := <-ch
		if ok {
			fmt.Printf("消费者得到数据 id: %d, recv: %d\n", id, value)
		} else {
			fmt.Printf("id: %d, closed\n", id)
			break
		}
	}
	done <- true
}

func TestProConDemo3(t *testing.T) {
	ch := make(chan int, 3)
	workerNum := 2
	done := make(chan bool, workerNum)
	for i := 1; i <= workerNum; i++ {
		go Consumer(i, ch, done)
	}
	go Producer(ch)
	for i := 1; i <= workerNum; i++ {
		// 由于生产者和消费者是异步，这里需要阻塞等待信号
		<-done
	}
}
