package codetop2023

// https://leetcode.cn/problems/implement-rand10-using-rand7/solution/xiang-xi-si-lu-ji-you-hua-si-lu-fen-xi-zhu-xing-ji/
func rand10() int {
	num := (rand7()-1)*7 + rand7()
	for num > 40 {
		num = (rand7()-1)*7 + rand7()
	}
	num = 1 + num%10
	return num
}

func rand7() int {
	return 0
}
