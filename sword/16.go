package sword

/**
*  n为偶数: y = x^ n/2, result = y*y
*  n为奇数: y = x^ n/2, result = y*y*x
*  x^n -> x^n/2 -> x^n/4 ->.. x^1 或 x^0
*  N为幂
 */

func myPow(x float64, n int) float64 {
	N := n
	if N > 0 {
		return quickPow(x, N)
	} else {
		return 1.0 / quickPow(x, -N)
	}
}

func quickPow(x float64, N int) float64 {
	if N == 0 {
		return 1.0
	}
	y := quickPow(x, N/2)
	if N%2 == 0 {
		return y * y
	} else {
		return y * y * x
	}
}
