package interview

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Test_SingleSender(t *testing.T) {
	// 单生产者，多消费者
	rand.Seed(time.Now().UnixNano())

	const MaxRandNumber = 100
	const MaxConsumer = 100

	wgConsumer := sync.WaitGroup{}
	wgConsumer.Add(MaxConsumer)

	dataCh := make(chan int, 100)

	go func() {
		for {
			if value := rand.Intn(MaxRandNumber); value == 0 {
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	for i := 0; i < MaxConsumer; i++ {
		go func() {
			defer wgConsumer.Done()
			for value := range dataCh {
				fmt.Println("value: ", value)
			}
		}()
	}

	wgConsumer.Wait()

}
