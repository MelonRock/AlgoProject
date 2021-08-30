package array

import (
	"math/rand"
)

type Solution struct {
	M        int
	BlackMap map[int]int
}

func Constructor710(n int, blacklist []int) Solution {
	blackMap := map[int]int{}
	for i := 0; i < len(blacklist); i++ {
		blackMap[blacklist[i]] = 1
	}
	M := n - len(blacklist)
	for _, value := range blacklist {
		if value < M {
			for {
				// 末尾开始映射。从 (M,N) 这个区间的末尾开始往前找，
				//找黑名单不存在的数，找到了就把 [0,M] 区间内冲突的数字映射到这里来
				if _, ok := blackMap[n-1]; ok {
					n--
				} else {
					break
				}
			}
			blackMap[value] = n - 1
			n--
		}
	}
	return Solution{BlackMap: blackMap, M: M}
}

func (this *Solution) Pick() int {
	idx := rand.Intn(this.M)
	if _, ok := this.BlackMap[idx]; ok {
		return this.BlackMap[idx]
	}
	return idx
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
