package interview

import (
	"sync"
	"sync/atomic"
	"testing"
)

// https://chai2010.cn/advanced-go-programming-book/ch1-basic/ch1-05-mem.html
func Test_LockSingleton(t *testing.T) {
	type singleton struct{}
	var (
		instance    *singleton
		initialized uint32
		mu          sync.Mutex
	)

	Instance := func() *singleton {
		if atomic.LoadUint32(&initialized) == 1 {
			return instance
		}

		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			defer atomic.StoreUint32(&initialized, 1)
			instance = &singleton{}
		}
		return instance
	}
	Instance()
}

func Test_OnceSingleton(t *testing.T) {
	type singleton struct{}
	var (
		instance *singleton
		once     sync.Once
	)
	Instance := func() *singleton {
		once.Do(func() {
			instance = &singleton{}
		})
		return instance
	}
	Instance()
}
