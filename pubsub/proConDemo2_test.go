package main

import (
	"fmt"
	"testing"
)

func producerCache(in chan<- int) {
	for i := 0; i < 10; i++ {
		data := i * i
		fmt.Println("生产者生产数据：", data)
		in <- data
	}
	close(in)
}

func consumerCache(out <-chan int) {
	for data := range out {
		fmt.Println("消费者得到数据:", data)
	}
}

func TestProConDemo2(t *testing.T) {
	ch := make(chan int, 5)
	go producerCache(ch)
	consumerCache(ch)
}
