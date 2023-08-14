package bitops

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	var res []byte
	for i >= 0 || j >= 0 {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
			j--
		}
		n := n1 + n2 + carry
		carry = n / 10
		res = append(res, byte(n%10)+'0')
	}
	if carry > 0 {
		res = append(res, '1')
	}
	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}
