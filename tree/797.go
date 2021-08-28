package tree

//
//
//// ***复习***
//func allPathsSourceTarget(graph [][]int) [][]int {
//	var (
//		res = [][]int{}
//		n = len(graph)
//		path = []int{}
//	)
//
//	var traverse3 func(int, []int)
//	traverse3 = func(i int, int []int) {
//		path = append(path, i)
//		if i == n-1{
//			res = append(res, append([]int{}, path)...)
//			path = path[:len(path)-1]
//			return
//		}
//		for _, node := range graph[i]{
//			traverse3(node, path)
//		}
//		path = path[:len(path)-1]
//	}
//	traverse3(0, path)
//	return res
//}
//
//
//
//
