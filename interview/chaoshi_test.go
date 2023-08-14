package interview

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("Hello")
		case <-ctx.Done():
			fmt.Println("Time Out")
			return ctx.Err()
		}
	}
}

func Test_Chaoshi(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}
	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}
