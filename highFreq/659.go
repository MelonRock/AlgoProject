package highFreq

func isPossible(nums []int) bool {
	n := len(nums)
	cmap, tmap := map[int]int{}, map[int]int{}
	for i := 0; i < n; i++ {
		cmap[nums[i]]++
		tmap[nums[i]] = 0
	}
	for i := 0; i < n; i++ {
		if cmap[nums[i]] > 0 {
			if tmap[nums[i]-1] > 0 {
				cmap[nums[i]]--
				tmap[nums[i]-1]--
				tmap[nums[i]]++
			} else {
				if cmap[nums[i]+1] > 0 && cmap[nums[i]+2] > 0 {
					cmap[nums[i]]--
					cmap[nums[i]+1]--
					cmap[nums[i]+2]--
					tmap[nums[i]+2]++
				} else {
					return false
				}
			}
		}
	}
	return true
}
