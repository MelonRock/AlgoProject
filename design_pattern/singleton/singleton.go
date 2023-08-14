package singleton

import (
	"sync"
)

type Singleton struct {
	a string
}

var lazySingleton *Singleton
var once = &sync.Once{}

func GetLazyInstance() *Singleton {
	if lazySingleton != nil {
		return lazySingleton
	}
	once.Do(func() {
		lazySingleton = &Singleton{a: "test"}
	})
	return lazySingleton
}
