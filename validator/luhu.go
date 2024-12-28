package validator

func ValidateNumber(cardNo string) bool {

	n := len(cardNo)
	var sum int
	isSecond := false
	for i := n - 1; i >= 0; i-- {
		d := cardNo[i]
		num := int(d - '0')
		if isSecond {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}

		sum += num
		isSecond = !isSecond

	}
	return (sum%10 == 0)

}
