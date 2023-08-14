package hot100

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		// 没重叠，直接合并
		if prev[1] < cur[0] {
			res = append(res, prev)
			prev = cur
		} else {
			// 合并
			prev[1] = max56(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

func max56(a, b int) int {
	if a > b {
		return a
	}
	return b
}
