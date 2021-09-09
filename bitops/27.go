package bitops

func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	slow := 0
	fast := 0
	// 快慢指针，每次遇到重复的就让fast前进，然后维护一个没有重复元素的nums[0..slow]
	for fast < n {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
