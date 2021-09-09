package bitops

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	slow := 0
	fast := 1
	// 快慢指针，每次遇到不重复的就让slow前进，然后维护一个没有重复元素的nums[0..slow]
	for fast < n {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
