package codetop2023

func mySqrt(x int) int {
	l, r := 1, x
	for l <= r {
		// 位操作，除以2并向下取整
		mid := (l + r) >> 1
		if x < mid*mid {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
