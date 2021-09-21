package sword

func permutation(s string) []string {
	res := []string{}
	b := []byte(s)
	var dfs38 func(x int)
	dfs38 = func(x int) {
		if x == len(b)-1 {
			res = append(res, string(b))
		}
		dict := make(map[byte]bool)
		for i := x; i < len(b); i++ {
			if !dict[b[i]] {
				b[x], b[i] = b[i], b[x]
				dict[b[x]] = true
				dfs38(x + 1)
				b[x], b[i] = b[i], b[x]
			}
		}
	}
	dfs38(0)
	return res
}
