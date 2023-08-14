package bytedance

import (
	"fmt"
	"testing"
	"time"
)

func TestBibao(t *testing.T) {
	values := []string{"a", "b", "c"}
	for _, val := range values {
		v := val
		go func() {
			fmt.Println(v)
		}()
	}

	values = []string{"a", "b", "c"}
	for _, val := range values {
		go func(v string) {
			fmt.Println(v)
		}(val)
	}
	time.Sleep(time.Second)
}

// 逃逸分析
