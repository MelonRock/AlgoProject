package codetop2023

import (
	"strconv"
	"strings"
)

var (
	path []string
	res  []string
)

func restoreIpAddresses(s string) []string {
	path, res = make([]string, 0, len(s)), make([]string, 0)
	dfs93(s, 0)
	return res
}

func dfs93(s string, start int) {
	// 终止条件为4段
	if len(path) == 4 {
		// 且遍历到最后一个字符
		if start == len(s) {
			str := strings.Join(path, ".")
			res = append(res, str)
		}
		return
	}

	for i := start; i < len(s); i++ {
		// 含有前导0的无效
		if i != start && s[start] == '0' {
			break
		}
		str := s[start : i+1]
		num, _ := strconv.Atoi(str)
		if num >= 0 && num <= 255 {
			// 符合条件的就进入下一层
			path = append(path, str)
			dfs93(s, i+1)
			path = path[:len(path)-1]
		} else {
			// 如果不满足条件，再往后也不可能满足条件，直接退出
			break
		}
	}
}
