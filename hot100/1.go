package hot100

func twoSum(nums []int, target int) []int {
	diffMap := map[int]int{}
	for i, v := range nums {
		if p, ok := diffMap[target-v]; ok {
			return []int{p, i}
		}
		diffMap[v] = i
	}
	return nil
}
