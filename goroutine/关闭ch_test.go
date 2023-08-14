package goroutine

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test_CloseCh(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var Max int = 10000
	var NumReceivers int = 10
	var NumSenders int = 1000

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	toStopCh := make(chan string, 1)
	//	var stoppedBy string
	go func() {
		_ = <-toStopCh
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					select {
					case toStopCh <- "sender#" + id:
					default:
					}
					return
				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			for {
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						select {
						case toStopCh <- "receiver#" + id:
						default:
						}
						return
					}
					fmt.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	select {
	case <-time.After(time.Hour):
	}

}
