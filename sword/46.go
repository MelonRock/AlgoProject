package sword

import "strconv"

func translateNum(num int) int {
	num_str := strconv.Itoa(num)
	res := 1
	p := 0
	q := 1
	for i := 1; i < len(num_str); i++ {
		p = q
		q = res
		sum := num_str[i-1 : i+1]
		if sum >= "10" && sum <= "25" {
			res += p
		}
	}
	return res
}
