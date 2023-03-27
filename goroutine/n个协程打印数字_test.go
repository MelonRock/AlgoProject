package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestNroutine(t *testing.T) {
	n := 5
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 1; j <= 100; j++ {
				if j%n == id {
					fmt.Printf("协程%d: %d\n", id, j)
				}
			}
		}(i)
	}
	wg.Wait()
}

func TestNroutineCh(t *testing.T) {
	gorutinenum := 5
	var chanslice []chan int
	exitchan := make(chan int)

	for i := 0; i < gorutinenum; i++ {
		chanslice = append(chanslice, make(chan int, 1))
	}

	res := 1
	j := 0
	for i := 0; i < gorutinenum; i++ {
		go func(i int) {
			for {
				<-chanslice[i]
				if res > 100 {
					exitchan <- 1
					break
				}

				fmt.Println(fmt.Sprintf("gorutine%v:%v", i, res))
				res++

				if j == gorutinenum-1 {
					j = 0
				} else {
					j++
				}
				chanslice[j] <- 1
			}
		}(i)
	}

	chanslice[0] <- 1

	select {
	case <-exitchan:
		fmt.Println("end")
	}
}
