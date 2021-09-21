package sword

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	l := 0
	r := len(matrix[0]) - 1
	t := 0
	b := len(matrix) - 1
	x := 0
	res := make([]int, (r+1)*(b+1))
	for {
		for i := l; i <= r; i++ {
			res[x] = matrix[t][i]
			x++
		}
		if t+1 > b {
			break
		}
		for i := t; i <= b; i++ {
			res[x] = matrix[i][r]
			x++
		}
		if l > r-1 {
			break
		}
		for i := r; i >= l; i-- {
			res[x] = matrix[b][i]
		}
		if t > b-1 {
			break
		}
		if l+1 > r {
			break
		}
	}
	return res
}
