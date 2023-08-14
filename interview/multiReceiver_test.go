package interview

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func Test_MultiReceiver(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	const MaxSender = 100
	const MaxReceiver = 100
	const MaxRandNumber = 1000

	wg := sync.WaitGroup{}
	wg.Add(MaxReceiver)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	notifyCh := make(chan string, 1)

	go func() {
		<-notifyCh
		close(stopCh)
	}()

	for i := 0; i < MaxSender; i++ {
		go func(i int) {
			for {
				value := rand.Intn(MaxRandNumber)
				if value == 0 {
					notifyCh <- "sender" + strconv.Itoa(i)
					return
				}

				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}

		}(i)
	}

	for i := 0; i < MaxReceiver; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == MaxRandNumber-1 {
						notifyCh <- "sender" + strconv.Itoa(i)
						return
					}
					fmt.Println("value", value)
				}
			}
		}(i)
	}

	wg.Wait()
}
