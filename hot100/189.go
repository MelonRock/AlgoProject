package hot100

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse189(nums)
	reverse189(nums[:k])
	reverse189(nums[k:])
}

func reverse189(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// 第i个数，移动到 i+k % n位置
//func rotate(nums []int, k int) {
//	n := len(nums)
//	temp := make([]int, n)
//	for i := 0; i < n; i++ {
//		temp[(i+k)%n] = nums[i]
//	}
//	for i := 0; i < n; i++ {
//		nums[i] = temp[i]
//	}
//}
