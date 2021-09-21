package sword

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	dfs34(root, target, []int{}, &res)
	return res
}

func dfs34(root *TreeNode, sum int, arr []int, res *[][]int) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)
	sum -= root.Val
	if sum == 0 && root.Left == nil && root.Right == nil {
		path := make([]int, len(arr))
		copy(path, arr)
		*res = append(*res, path)
	}
	dfs34(root.Left, sum, arr, res)
	dfs34(root.Right, sum, arr, res)
	arr = arr[:len(arr)-1]
}
