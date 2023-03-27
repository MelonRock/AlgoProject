package goroutine

import (
	"fmt"
	"sync"
)

func main() {
	inChan := make(chan int)
	outChan := make(chan int)
	var wg sync.WaitGroup

	// Start two producers
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				inChan <- i*10 + j
			}
		}(i)
	}

	// Start three consumers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range inChan {
				outChan <- val * 2
			}
		}()
	}

	// Wait for all producers to finish
	go func() {
		wg.Wait()
		close(inChan)
	}()

	// Wait for all consumers to finish
	go func() {
		wg.Wait()
		close(outChan)
	}()

	// Print all results from the outChan
	for val := range outChan {
		fmt.Println(val)
	}
}
