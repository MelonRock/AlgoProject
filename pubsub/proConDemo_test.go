package main

import (
	"fmt"
	"testing"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		data := i * i
		fmt.Println("生产者生产数据：", data)
		out <- data
	}
	close(out)
}

func consumer(in <-chan int) {
	for data := range in {
		fmt.Println("消费者得到数据：", data)
	}
}

func TestProConDemo(t *testing.T) {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
