package codetop

type ListNode1 struct {
	key, val  int
	pre, next *ListNode1
}

type LRUCache struct {
	m          map[int]*ListNode1
	head, tail *ListNode1
	capacity   int
}

func Constructor(capacity int) LRUCache {
	head := &ListNode1{pre: nil, next: nil}
	tail := &ListNode1{pre: nil, next: nil}
	head.next = tail
	tail.pre = head
	return LRUCache{
		m:        make(map[int]*ListNode1),
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; !ok {
		return -1
	} else {
		this.MoveToHead(node)
		return node.key
	}
}

func (this *LRUCache) Put(key int, value int) {
	tail := this.tail
	cache := this.m
	if node, ok := cache[key]; ok {
		// 存在，更新值
		node.val = value
		cache[key] = node
		this.MoveToHead(node)
	} else {
		// 不存在，队列满和不满
		v := &ListNode1{key: key, val: value, pre: nil, next: nil}
		if len(cache) == this.capacity {
			// 队列满了，删除最后一个
			delete(cache, tail.pre.key)
			this.DeleteNode(tail.pre)
		}
		this.AddNode(v)
		cache[key] = v
	}
}

func (this *LRUCache) MoveToHead(node *ListNode1) {
	this.DeleteNode(node)
	this.AddNode(node)
}

func (this *LRUCache) DeleteNode(node *ListNode1) {
	node.pre.next = node.pre
	node.next.pre = node.pre
}

func (this *LRUCache) AddNode(node *ListNode1) {
	head := this.head
	next := head.next
	head.next = node
	node.pre = head
	node.next = next
	next.pre = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
