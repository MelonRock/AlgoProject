package codetop

func sortArray(nums []int) []int {
	quickSort1(nums, 0, len(nums)-1)
	return nums
}

func quickSort1(num []int, left, right int) {
	if left > right {
		return
	}

	i, j, pivot := left, right, num[left]
	for i < j {
		for num[i] <= pivot && i < j {
			i++
		}

		for num[j] >= pivot && i < j {
			j--
		}
		num[i], num[j] = num[j], num[i]
	}
	num[i], num[left] = num[left], num[j]
	quickSort1(num, left, i-1)
	quickSort1(num, i+1, right)
}
