package sword

// https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/solution/mian-shi-ti-10-ii-qing-wa-tiao-tai-jie-wen-ti-dong/

func numWays(n int) int {
	prev, curr := 1, 1
	for i := 2; i <= n; i++ {
		next := (prev + curr) % 1000000007
		prev = curr
		curr = next
	}
	return curr
}
