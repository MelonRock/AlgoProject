package bitops

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	for _, x := range bookings {
		nums[x[0]-1] += x[2]
		if x[1] < n {
			nums[x[1]] -= x[2]
		}
	}
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
