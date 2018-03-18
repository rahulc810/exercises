package easy

import "fmt"

func ConvertHours(hh int) (int, string) {
	fmt.Printf("ConvertHours called with value %d\n", hh)
	if hh > 24 || hh < 0 {
		panic("invalid hours field")
	}
	adjustedHour := hh % 12
	var phase string
	if hh/12 == 0 {
		phase = "AM"
	} else {
		phase = "PM"
	}

	return adjustedHour, phase
}

func TranslateToWords(num int, wordZero bool) string {
	fmt.Printf("ConvertWords called with value %d\n", num)
	if num < 0 || num > 99 {
		panic("Invalid minute value")
	}
	var ret string
	if isTeen(num) {
		ret = translateTeens(num)
	} else {
		unitPlaceDigit := num % 10
		tenPlaceDigit := num / 10
		x := translateUnits(unitPlaceDigit)
		var y string
		if wordZero {
			y = translateTens(tenPlaceDigit)
		} else if tenPlaceDigit != 0 {
			y = translateTens(tenPlaceDigit)
		}
		ret = y + " " + x
	}
	return ret
}

func translateUnits(digit int) string {
	if digit < 0 || digit > 9 {
		panic(fmt.Sprintf("Illegal digit in units %d \n ", digit))
	}
	var ret string
	switch digit {
	case 0:
		ret = "zero"
	case 1:
		ret = "one"
	case 2:
		ret = "two"
	case 3:
		ret = "three"
	case 4:
		ret = "four"
	case 5:
		ret = "five"
	case 6:
		ret = "six"
	case 7:
		ret = "seven"
	case 8:
		ret = "eight"
	case 9:
		ret = "nine"
	}
	return ret
}

func translateTens(digit int) string {
	if digit < 0 || digit > 9 {
		panic(fmt.Sprintf("Illegal digit in tens %d \n ", digit))
	}
	var ret string
	switch digit {
	case 0:
		ret = "zero"
	case 1:
		ret = "-"
	case 2:
		ret = "twenty"
	case 3:
		ret = "thirty"
	case 4:
		ret = "forty"
	case 5:
		ret = "fifty"
	case 6:
		ret = "sixty"
	case 7:
		ret = "seventy"
	case 8:
		ret = "eighty"
	case 9:
		ret = "ninety"
	}
	return ret
}

func translateTeens(digit int) string {
	if digit < 10 || digit > 19 {
		panic("Illegal digit")
	}
	var ret string
	switch digit {
	case 10:
		ret = "Ten"
	case 11:
		ret = "Eleven"
	case 12:
		ret = "Twelve"
	case 13:
		ret = "Thirteen"
	case 14:
		ret = "Fourteen"
	case 15:
		ret = "Fifteen"
	case 16:
		ret = "Sixteen"
	case 17:
		ret = "Seventeen"
	case 18:
		ret = "Eighteen"
	case 19:
		ret = "Nineteen"
	}
	return ret
}

func isTeen(mm int) bool {
	return mm < 20 && mm > 9
}
