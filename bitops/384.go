package bitops

import "math/rand"

type Solution struct {
	arr []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.arr
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := len(this.arr)
	res := make([]int, n)
	copy(res, this.arr)
	for i := n - 1; i >= 0; i-- {
		rand := rand.Intn(i + 1) // math.rand中的Intn(i+1)返回[0, i]范围的整数
		res[i], res[rand] = res[rand], res[i]
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
