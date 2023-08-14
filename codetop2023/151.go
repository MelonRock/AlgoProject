package codetop2023

import "strings"

func reverseWords(s string) string {
	// 删除首位空格
	s = strings.TrimSpace(s)
	b := []byte(s)
	i, j := len(b)-1, len(b)-1
	reverse := make([]string, 0)
	for i >= 0 {
		// 搜索首个空格
		for i >= 0 && b[i] != ' ' {
			i--
		}
		// 添加单词
		reverse = append(reverse, string(b[i+1:j+1]))
		// 跳过单词间空格
		for i >= 0 && b[i] == ' ' {
			i--
		}
		// 指向下个单词的尾字符
		j = i
	}
	res := strings.Join(reverse, " ")
	return res
}
