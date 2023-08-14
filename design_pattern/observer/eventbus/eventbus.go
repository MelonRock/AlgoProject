package eventbus

import (
	"fmt"
	"reflect"
	"sync"
)

type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

type AsyncEventBus struct {
	handler map[string][]reflect.Value
	lock    sync.Mutex
}

func NewAsynEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handler: map[string][]reflect.Value{},
		lock:    sync.Mutex{},
	}
}

func (bus *AsyncEventBus) Subscribe(topic string, f interface{}) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	v := reflect.ValueOf(f)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler is not a function")
	}

	handler, ok := bus.handler[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, v)
	bus.handler[topic] = handler
	return nil

}

func (bus *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := bus.handler[topic]
	if !ok {
		fmt.Printf("not found handlers in topic:%s", topic)
		return
	}
	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}
	for i := range handlers {
		go handlers[i].Call(params)
	}
}
