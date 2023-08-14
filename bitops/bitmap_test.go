package bitops

import (
	"fmt"
	"testing"
)

func index(num int) int {
	// 映射函数，将整数映射到字符中的位置
	return num / 8 // 每个字符存储8个整数
}

func shift(num int) int {
	// 转换函数，将整数转换为字符中的位
	return num % 8 // 整数除以8的余数表示在字符中的位数
}

func bitmapSort(array []int, byteArray []byte, size int) []int {
	for i := 0; i < size; i++ {
		j := index(array[i])
		k := shift(array[i])
		byteArray[j] |= byte(k>>0x1f) & 0x1 << byte(k&0x7) // 对byteArray数组中的字符执行按位或（OR）操作
	}
	var index int
	for i := 0; i < len(byteArray); i++ {
		for j := 0; j < 8; j++ {
			if (byteArray[i]>>byte(j))&0x1 == 1 { // 遍历字符数组，如果某个位置的值为1，则将对应的整数存入新的整型数组中
				array[index] = i*8 + j
				index++
			}
		}
	}
	return array
}

func Test_Bitmap(t *testing.T) {
	var k int
	k = 2
	fmt.Println(byte(k >> 0x1f))
	fmt.Println(0x7)
}
