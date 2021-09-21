package sword

func validateStackSequences(push []int, popped []int) bool {
	stack := make([]int, 0)
	i := 0
	for _, value := range push {
		stack = append(stack, value)
		for len(stack) != 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
