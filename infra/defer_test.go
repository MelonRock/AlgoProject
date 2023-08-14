package infra

import (
	"fmt"
	"testing"
)

func Test_Defer(t *testing.T) {
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())
	fmt.Println(test4())

	return
}

func test1() (v int) {
	defer fmt.Println("defer test1", v)
	return v
}

func test2() (v int) {
	defer func() {
		fmt.Println("defer test2", v)
	}()
	return 3
}

func test3() (v int) {
	defer fmt.Println("defer test3", v)
	v = 3
	return 4
}

func test4() (v int) {
	defer func(n int) {
		fmt.Println("defer test4", n)
	}(v)
	return 5
}
