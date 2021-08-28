package tree

/**
* Definition for a Node.
* type Node struct {
*     Val int
*     Left *Node
*     Right *Node
*     Next *Node
* }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(left *Node, right *Node) {
	if left == nil || right == nil {
		return
	}
	// 前序遍历
	left.Next = right
	// 左右节点在同一棵树
	connectTwoNode(left.Left, left.Right)
	connectTwoNode(right.Left, right.Right)
	// 左右节点不在同一棵树
	connectTwoNode(left.Right, right.Left)
}
