package codetop2023

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for index, value := range nums {
		if p, ok := hashTable[target-value]; ok {
			return []int{p, index}
		}
		hashTable[value] = index
	}
	return []int{}
}
