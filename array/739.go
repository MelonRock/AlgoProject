package array

/*
给定一个整数数组temperatures，表示每天的温度，
返回一个数组answer，其中answer[i]是指对于第 i 天，
下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用0 来代替。

*/
func dailyTemperatures(t []int) []int {
	var res []int
	for i := 0; i < len(t)-1; i++ {
		j := i + 1
		for ; j < len(t); j++ {
			// 如果之后出现更高，说明找到了
			if t[j] > t[i] {
				res = append(res, j-i)
				break
			}
		}
		// 找完了都没找到
		if j == len(t) {
			res = append(res, 0)
		}
	}
	// 最后一个肯定是 0
	return append(res, 0)
}
