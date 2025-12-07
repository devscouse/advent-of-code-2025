package common

import "math"

func CountDigits(value int) int {
	if value < 0 {
		value = value * -1
	}
	comparison := 10
	ndigits := 1
	for {
		if value < comparison {
			return ndigits
		}
		comparison *= 10
		ndigits++
	}
}

func TakeLeastSignificantDigit(number *int) int {
	if *number < 10 {
		digit := *number
		*number = 0
		return digit
	}
	digit := *number % 10
	*number = *number / 10
	return digit
}

func TakeMostSignificantDigit(number int) (int, int) {
	if number < 10 {
		digit := number
		number = 0
		return 0, digit
	}
	factor := int(math.Pow10(CountDigits(number)))
	digit := number / factor
	number = number % factor
	return number, digit
}

func AddLeastSignificantDigit(number int, digit int) int {
	return number*10 + digit
}
