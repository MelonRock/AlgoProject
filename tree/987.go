package tree

import "sort"

// 存放坐标与节点
type point struct {
	t   *TreeNode
	row int
	col int
}

// 层级kv内容
type screening struct {
	key   int
	value int
}

func verticalTraversal(root *TreeNode) [][]int {
	mp := make(map[int][]int)
	sequence := []point{{
		t:   root,
		row: 0,
		col: 0,
	}}
	// 层序遍历
	appendScreening := make([]screening, 0)
	for len(sequence) != 0 {
		ln := len(sequence)
		for i := 0; i < ln; i++ {
			v := sequence[i]
			// 放入层序内容，因为答案需要层级有序，所以后面要排序
			appendScreening = append(appendScreening, screening{
				key:   v.col,
				value: v.t.Val,
			})
			if v.t.Left != nil {
				sequence = append(sequence, point{
					t:   v.t.Left,
					row: v.row + 1,
					col: v.col - 1,
				})
			}
			if v.t.Right != nil {
				sequence = append(sequence, point{
					t:   v.t.Right,
					row: v.row + 1,
					col: v.col + 1,
				})
			}
		}
		// 简单排序，主要排value, key不需要，使用map来筛选
		sort.Slice(appendScreening, func(i, j int) bool {
			return appendScreening[i].value < appendScreening[j].value
		})
		// 分别加到答案mp中
		for _, v := range appendScreening {
			mp[v.key] = append(mp[v.key], v.value)
		}
		// 置空变量
		appendScreening = appendScreening[:0]
		// 弹出元素
		sequence = sequence[ln:]
	}

	res := make([][]int, 0, len(mp))
	ks := make([]int, 0, len(mp))
	// 将答案key取出
	for k := range mp {
		ks = append(ks, k)
	}
	// 排序key
	sort.Ints(ks)
	// 根据key取出答案
	for _, v := range ks {
		res = append(res, mp[v])
	}
	return res
}
