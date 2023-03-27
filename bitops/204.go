package bitops

func countPrimes(n int) int {
	isNotPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if isNotPrime[i] {
			continue
		}
		// 每计算一个数，都要把它的倍数去掉。到了n，数一下留下了几个数。
		for j := i * i; j < n; j = j + i {
			isNotPrime[j] = true
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			count++
		}
	}
	return count
}
