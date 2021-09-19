package highFreq

func pancakeSort(arr []int) []int {
	n := len(arr)
	res := []int{}
	var reverseArray func(right int)
	reverseArray = func(right int) {
		left := 0
		for left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	var sort func(size int)
	sort = func(size int) {
		if size <= 1 {
			return
		}
		maxPos := 0
		max := arr[0]
		for i := 0; i < size; i++ {
			if arr[i] > max {
				maxPos = i
				max = arr[i]
			}
		}
		reverseArray(maxPos)
		res = append(res, maxPos+1)
		reverseArray(size - 1)
		res = append(res, size)
		sort(size - 1)
	}
	sort(n)
	return res
}
