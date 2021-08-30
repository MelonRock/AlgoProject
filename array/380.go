package array

import "math/rand"

type RandomizedSet struct {
	sl []int
	m  map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{m: map[int]int{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.sl = append(this.sl, val)
	this.m[val] = len(this.sl) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	curIdx := this.m[val]
	lastIdx := len(this.sl) - 1
	lastVal := this.sl[lastIdx]
	// 交换删除元素和末尾元素的位置
	this.sl[curIdx], this.sl[lastIdx] = lastVal, val
	// 直接pop末尾元素， O（1）时间复杂度
	this.sl = this.sl[:lastIdx]
	// 更新map的被交换元素index值
	this.m[lastVal] = curIdx
	delete(this.m, val) // map删除元素
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.sl[rand.Int()%len(this.sl)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
