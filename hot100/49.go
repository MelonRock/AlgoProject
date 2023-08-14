package hot100

import "sort"

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for i, _ := range strs {
		s := []byte(strs[i])
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], strs[i])
	}
	res := make([][]string, 0)
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}
