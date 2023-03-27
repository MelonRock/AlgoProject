package codetop2023

import "sort"

func merge56(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] {
			// 两个区间没有重合
			res = append(res, prev)
			prev = cur
		} else {
			// 有重合部分，需要合并找出较大的右边界
			prev[1] = max(prev[1], cur[1])
		}
	}
	// 合并最后一个区间
	res = append(res, prev)
	return res
}
