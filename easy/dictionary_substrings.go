package easy

type coord struct {
	s, e             int
	largestConsonant rune
}

var MaxBk coord
var MinBk coord

var alpha = map[rune][]rune{
	'a': []rune{'b', 'c', 'd'},
	'e': []rune{'f', 'g', 'h'},
	'i': []rune{'j', 'k', 'l', 'm', 'n'},
	'o': []rune{'p', 'q', 'r', 's', 't'},
	'u': []rune{'v', 'w', 'x', 'y', 'z'},
}

func FindDctSubString(s []rune) {
	max := coord{-1, -1, '0'}
	min := coord{-1, -1, '0'}

	for idx, char := range s {

	}
}

func adjustMin(s []rune, min coord, key int) {
	if isVowel(s[key]) {
		if s[min.s] > s[key] {
			updateMinBk(s, min)
			min.s = key
			min.e = -1
		} else {

		}
	} else {
		if min.largestConsonant == '0' {
			min.e = key
			min.largestConsonant = s[key]
			updateMinBk(s, min)
		}
	}

}

func updateMinBk(s []rune, min coord) {
	if isValidCoord(min) {
		if isValidCoord(MinBk) {
			tempBk := coord{min.s, min.e, min.largestConsonant}
			out := compare(s, MinBk, tempBk)
			if out > 1 {
				MinBk = tempBk
			}
		}
	}
}

func adjustMax(s []rune, max coord, key int) {
	if isVowel(s[key]) {
		if s[max.s] < s[key] {
			if isValidCoord(max) {
				MaxBk = coord{max.s, max.e, max.largestConsonant}
			}
			max.s = key
		}
	} else {
		max.e = key
	}
}

func compare(s []rune, a, b coord) int {
	i, j := 0, 0
	for i, j = a.s, b.s; i <= a.e && j <= b.e; i, j = i+1, j+1 {
		if s[i] > s[j] {
			return 1
		} else if s[i] < s[j] {
			return -1
		} else {
			continue
		}
	}
	if i < j {
		return -1
	} else if i > j {
		return 1
	} else {
		return 0
	}
}

func isValidCoord(c coord) bool {
	return c.s > -1 && c.e > -1
}

func isVowel(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
