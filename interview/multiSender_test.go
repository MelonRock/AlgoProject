package interview

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Test_MultiSender(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	const MaxRandNumber = 100
	const MaxSender = 100

	wg := sync.WaitGroup{}
	wg.Add(1)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	for i := 0; i < MaxSender; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(MaxRandNumber):
				}
			}
		}()
	}

	go func() {
		defer wg.Done()
		for value := range dataCh {
			if value == MaxRandNumber-1 {
				close(stopCh)
				return
			}
			fmt.Println("value: ", value)
		}
	}()

	wg.Wait()

}
