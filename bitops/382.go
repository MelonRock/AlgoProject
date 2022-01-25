package bitops

import "math/rand"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution1 struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor1(head *ListNode) Solution1 {
	return Solution1{head: head}
}

/** Returns a random node's value. */
func (this *Solution1) GetRandom() int {
	var ans, count int
	cur := this.head
	for cur != nil {
		count++
		randCount := rand.Int()%count + 1
		if randCount == count {
			ans = cur.Val
		}
		cur = cur.Next
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
