package hot100

func calculate(s string) int {
	nums := []int{} // 栈
	curNum := 0
	preOp := '+' // 上一个遇到的运算符
	s = s + "+"  // 为了计算最后一个数

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			curNum = curNum*10 + int(s[i]-'0') // 进制转换
		} else if s[i] == ' ' {
			continue
		} else { // 遇到运算符
			if preOp == '+' { // 如果上一个遇到加号
				nums = append(nums, curNum)
			} else if preOp == '-' {
				nums = append(nums, -1*curNum)
			} else if preOp == '*' {
				nums[len(nums)-1] = nums[len(nums)-1] * curNum // 更新栈顶元素为乘过的值
			} else if preOp == '/' {
				nums[len(nums)-1] = nums[len(nums)-1] / curNum // 更新栈顶元素为除过的值
			}
			preOp = rune(s[i]) // 记录当前遇到的运算符
			curNum = 0         // 重置curnums
		}
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
