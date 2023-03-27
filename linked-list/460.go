package linked_list

import (
	"container/list"
)

type LFUCache struct {
	nodes    map[int]*list.Element // findMap
	lists    map[int]*list.List    // countMap
	capacity int
	min      int
}

type node struct {
	key       int
	value     int
	frequency int
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{
		nodes:    make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	// findMap中存在元素
	if currentValue, ok := this.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		// 更新value值
		currentNode.value = value
		// 增加访问频率
		this.Get(key)
		return
	}
	// key不存在，队列满了
	if this.capacity == len(this.nodes) {
		// 容量已满的话需要淘汰一个 freq 最小的 key
		currentList := this.lists[this.min]
		// 队首元素是最先入队的访问次数最少元素
		frontNode := currentList.Front()
		// 在findMap中移除
		delete(this.nodes, frontNode.Value.(*node).key)
		// 在countMap中移除
		currentList.Remove(frontNode)
	}
	// 新添加的元素，更新freq_min
	this.min = 1
	currentNode := &node{
		key:       key,
		value:     value,
		frequency: 1,
	}
	// 不存在其他使用次数为1的缓存
	if _, ok := this.lists[1]; !ok {
		this.lists[1] = list.New()
	}
	newList := this.lists[1]
	// 将 key 加入 freq = 1 对应的列表中
	newNode := newList.PushBack(currentNode)
	// 更新findMap中元素的位置
	this.nodes[key] = newNode
}

func (this *LFUCache) Get(key int) int {
	// findMap找到元素
	value, ok := this.nodes[key]
	if !ok {
		return -1
	}
	// 类型转换
	currentNode := value.Value.(*node)
	// 从countMap中移除该元素
	this.lists[currentNode.frequency].Remove(value)
	// 更新频率
	currentNode.frequency++
	// 不存在其他使用次数为freq+1的缓存
	if _, ok := this.lists[currentNode.frequency]; !ok {
		this.lists[currentNode.frequency] = list.New()
	}
	// 将 key 加入 freq + 1 对应的列表中
	newList := this.lists[currentNode.frequency]
	newNode := newList.PushBack(currentNode)
	// 更新findMap中元素的位置
	this.nodes[key] = newNode
	// 如果这个 freq 恰好是 minFreq，更新 minFreq. 且如果 freq -1 对应的列表空了
	if currentNode.frequency-1 == this.min && this.lists[currentNode.frequency-1].Len() == 0 {
		this.min++
	}
	return currentNode.value
}
