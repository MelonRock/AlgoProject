package tree

import (
	"fmt"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var memo = make(map[string]int)
	var res []*TreeNode
	traverse1(root, &res, &memo)
	return res
}

func traverse1(root *TreeNode, res *[]*TreeNode, memo *map[string]int) string {
	// base case
	if root == nil {
		return "#"
	}
	// 将左右子树序列化成字符串
	left := traverse1(root.Left, res, memo)
	right := traverse1(root.Right, res, memo)

	// 左右子树加上自己，就是以自己为根的二叉树序列化结果
	subtree := left + "," + right + "," + strconv.Itoa(root.Val)
	fmt.Print(subtree)
	// 未重复的，加入map中
	if freq, ok := (*memo)[subtree]; !ok {
		(*memo)[subtree] = 1
	} else {
		// 第一次重复的加入结果
		if freq == 1 {
			*res = append(*res, root)
		}
		// 子树对应出现的次数加1
		(*memo)[subtree] = freq + 1
	}
	return subtree
}
