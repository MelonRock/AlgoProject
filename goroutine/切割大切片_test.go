package goroutine

import (
	"fmt"
	"testing"
)

func splitSlice(slice []int, size int) [][]int {
	var result [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}

func TestSplitSlice(t *testing.T) {
	sLen := 11
	size := 5
	s := make([]int, 0, sLen)
	for i := 0; i < sLen; i++ {
		s = append(s, i)
	}
	result := splitSlice(s, size)
	for i, sp := range result {
		fmt.Println(i, " ==> ", sp)
	}
}
