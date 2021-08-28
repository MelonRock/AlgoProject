package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree2(inorder []int, postorder []int) *TreeNode {
	root := build2(inorder, 0, len(inorder)-1,
		postorder, 0, len(postorder)-1)
	return root
}

func build2(inorder []int, inStart int, inEnd int,
	postorder []int, postStart int, postEnd int) *TreeNode {
	// base case
	if inStart > inEnd {
		return nil
	}
	// postEnd为根节点
	rootVal := postorder[postEnd]
	// 找到在inorder的index
	index := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}
	// 找出inorder左子树的大小
	leftSize := index - inStart
	// 构造当前根节点
	root := &TreeNode{Val: rootVal}
	// 递归遍历左右子树
	root.Left = build2(inorder, inStart, index-1,
		postorder, postStart, postStart+leftSize-1)
	root.Right = build2(inorder, index+1, inEnd,
		postorder, postStart+leftSize, postEnd-1)
	return root
}
