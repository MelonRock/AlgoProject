package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

// 将 nums[lo..hi] 构造成符合条件的树，返回根节点 */
func build(nums []int, lo int, hi int) *TreeNode {
	// base case
	if lo > hi {
		return nil
	}
	// 找到数组中的最大值和对应的索引
	index := lo
	for i := lo + 1; i <= hi; i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}

	root := &TreeNode{Val: nums[index]}
	// 递归调用构造左右子树
	root.Left = build(nums, lo, index-1)
	root.Right = build(nums, index+1, hi)
	return root
}
