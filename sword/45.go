package sword

import (
	"fmt"
	"sort"
	"strings"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compareNumber(nums[i], nums[j])
	})
	var res strings.Builder
	for i := 0; i < len(nums); i++ {
		res.WriteString(fmt.Sprintf("%d", nums[i]))
	}
	return res.String()
}

func compareNumber(a, b int) bool {
	str1 := fmt.Sprint("%d%d", a, b)
	str2 := fmt.Sprint("%d%d", b, a)
	if str1 < str2 {
		return true
	}
	return false
}
