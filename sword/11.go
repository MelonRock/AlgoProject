package sword

func minArray(number []int) int {
	left := 0
	right := len(number) - 1

	for left < right {
		mid := left + (right-left)>>1
		if number[mid] > number[right] {
			left = mid + 1
		} else if number[mid] <= number[right] {
			right = right - 1
		}
	}
	return number[left]
}
