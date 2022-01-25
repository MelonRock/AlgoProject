package bitops

import (
	"runtime"
	"testing"
	"time"
)

func TestInterView(t *testing.T) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go println(i)
	}
	runtime.Gosched()
	time.Sleep(time.Second)
}
