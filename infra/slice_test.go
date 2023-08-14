package infra

import (
	"fmt"
	"testing"
)

func Test_Slice(t *testing.T) {
	arr := []int{1}

	myfunc1(arr)
	fmt.Printf("%p, len:%d, cap:%d, %v\n", arr, len(arr), cap(arr), arr)

	arr = append(arr, 3)
	arr = append(arr, 4) // [1 3 4]
	fmt.Printf("%p, len:%d, cap:%d, %v\n", arr, len(arr), cap(arr), arr)
	myfunc2(arr)
	fmt.Printf("%p, len:%d, cap:%d, %v\n", arr, len(arr), cap(arr), arr) // [9 3 4]
}

func myfunc1(arr []int) {
	arr = append(arr, 2) // 发生扩容
	arr[0] = 0
	fmt.Printf("%p, len:%d, cap:%d, %v\n", arr, len(arr), cap(arr), arr) // [0 2]
	return
}

func myfunc2(arr []int) {
	arr = append(arr, 5) // [1 3 4 5]
	arr[0] = 9
	fmt.Printf("%p, len:%d, cap:%d, %v\n", arr, len(arr), cap(arr), arr) // [9 3 4 5]
	return
}
