package hot100

func wordBreak(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	return dfs139(0, s, wordMap)
}

// 判断从start到末尾的子串能否break
func dfs139(start int, s string, wordMap map[string]bool) bool {
	// 指针越界，s一步步成功划分为单词，才走到越界这步，现在没有剩余子串
	if start == len(s) {
		return true
	}
	for i := start + 1; i <= len(s); i++ {
		prefix := s[start:i]
		// 如果前缀部分不是单词，就不会执行canBreak(i)。进入下一轮迭代，再切出一个前缀串，再试
		if wordMap[prefix] && dfs139(i, s, wordMap) {
			return true
		}
	}
	return false
}
