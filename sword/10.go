package sword

func numWays(n int) int {
	prev, curr := 1, 1
	for i := 2; i <= n; i++ {
		next := (prev + curr) % 1000000007
		prev = curr
		curr = next
	}
	return curr
}
