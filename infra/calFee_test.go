package infra

func calculateFee(orderCount int) float64 {
	var totalFee float64
	for orderCount > 0 {
		switch {
		case orderCount >= 51:
			totalFee += float64(orderCount-50) * 1.0
			orderCount = 50
		case orderCount >= 21:
			totalFee += float64(orderCount-20) * 10.0
			orderCount = 20
		case orderCount >= 6:
			totalFee += float64(orderCount-5) * 15.0
			orderCount = 5
		default:
			totalFee += float64(orderCount) * 30.0
			orderCount = 0
		}
	}
	return totalFee
}
