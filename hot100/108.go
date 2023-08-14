package hot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	midIdx := (start + end) >> 1
	root := &TreeNode{Val: nums[midIdx]}

	root.Left = buildBST(nums, start, midIdx-1)
	root.Right = buildBST(nums, midIdx+1, end)
	return root
}
