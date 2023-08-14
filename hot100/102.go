package hot100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		level := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[i]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[l:]
		res = append(res, level)
	}
	return res
}
