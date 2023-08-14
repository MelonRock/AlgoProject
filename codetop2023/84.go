package codetop2023

func largestRectangleArea(hs []int) int {
	n := len(hs)
	l := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = -1
		r[i] = n
	}

	d := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(d) > 0 && hs[d[len(d)-1]] > hs[i] {
			r[d[len(d)-1]] = i
			d = d[:len(d)-1]
		}
		d = append(d, i)
	}

	d = nil
	d = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(d) > 0 && hs[d[len(d)-1]] > hs[i] {
			l[d[len(d)-1]] = i
			d = d[:len(d)-1]
		}
		d = append(d, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		h := hs[i]
		a := l[i]
		b := r[i]
		ans = max(ans, h*(b-a-1))
	}
	return ans
}
