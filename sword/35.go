package sword

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head
	nodes := make(map[*Node]*Node)
	// 复制各节点，并建立 原节点->新节点的map映射
	for cur != nil {
		nodes[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	// 4. 构建新链表的 next 和 random 指向
	for cur != nil {
		nodes[cur].Next = nodes[cur.Next]
		nodes[cur].Random = nodes[cur.Random]
		cur = cur.Next
	}
	// 5. 返回新链表的头节点
	return nodes[head]
}
