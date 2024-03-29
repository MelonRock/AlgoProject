package eventbus

import (
	"fmt"
	"testing"
	"time"
)

func sub1(msg1, msg2 string) {
	time.Sleep(1 * time.Microsecond)
	fmt.Printf("sub1, %s %s \n", msg1, msg2)
}

func sub2(msg1, msg2 string) {
	fmt.Printf("sub2, %s %s \n", msg1, msg2)
}

func TestAsynEventBus_Publish(t *testing.T) {
	bus := NewAsynEventBus()
	bus.Subscribe("topic:1", sub1)
	bus.Subscribe("topic:2", sub2)
	bus.Publish("topic:1", "test1", "test2")
	bus.Publish("topic:2", "test1", "test2")
	time.Sleep(1 * time.Second)
}
