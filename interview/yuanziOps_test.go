package interview

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// 有锁
func Test_LockYuanzi(t *testing.T) {
	var total struct {
		sync.Mutex
		value int
	}

	worker := func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i <= 100; i++ {
			total.Lock()
			total.value += i
			total.Unlock()
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total.value)
}

func Test_AtomicYuanzi(t *testing.T) {
	var total uint64

	worker := func(wg *sync.WaitGroup) {
		defer wg.Done()
		var i uint64
		for i = 0; i <= 100; i++ {
			atomic.AddUint64(&total, i)
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total)
}
