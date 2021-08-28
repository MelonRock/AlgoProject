package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	root := build1(preorder, 0, len(preorder)-1,
		inorder, 0, len(inorder)-1)
	return root
}

func build1(preorder []int, preStart int, preEnd int,
	inorder []int, inStart int, inEnd int) *TreeNode {
	// base case
	if preStart > preEnd {
		return nil
	}
	// preStart为根节点
	rootVal := preorder[preStart]
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
	root.Left = build1(preorder, preStart+1, preStart+leftSize,
		inorder, inStart, index-1)
	root.Right = build1(preorder, preStart+leftSize+1, preEnd,
		inorder, index+1, inEnd)
	return root
}
