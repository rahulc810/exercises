package easy

import (
	"fmt"
)

type coord struct {
	s, e                           int
	posStartingVowelRecur          int
	largestConsonant, largestVowel int
}

var MaxBk *coord
var MinBk *coord

var alpha = map[rune][]rune{
	'a': []rune{'b', 'c', 'd'},
	'e': []rune{'f', 'g', 'h'},
	'i': []rune{'j', 'k', 'l', 'm', 'n'},
	'o': []rune{'p', 'q', 'r', 's', 't'},
	'u': []rune{'v', 'w', 'x', 'y', 'z'},
}

func FindDctSubString() {
	strs := [][]rune{[]rune("aaaaab"), []rune("eaaaaaab"), []rune("eoxuez"), []rune("eoueg")}
	//strs = [][]rune{[]rune("eoueg")}
	for _, s := range strs {
		max := &coord{-1, -1, -1, -1, -1}
		min := &coord{-1, -1, -1, -1, -1}
		for idx, _ := range s {
			adjustMin(s, min, idx)
			adjustMax(s, max, idx)
		}

		fmt.Printf("%q\n", s)

		fmt.Printf("MAX - ")
		printCoord(max, s)

		fmt.Printf("MIN - ")
		printCoord(min, s)

		fmt.Println()
	}
}

func printCoord(c *coord, s []rune) {
	for idx, _ := range s {
		if idx >= c.s && idx <= c.e {
			fmt.Printf("%q", s[idx])
		}
	}
	fmt.Println()
}

func adjustMin(s []rune, min *coord, key int) {
	if isVowel(s[key]) {
		if min.s < 0 || s[min.s] > s[key] {
			min.s = key
			min.e = -1
			min.largestVowel = key
		} else {
			if s[min.s] < s[key] && s[min.largestVowel] < s[key] {
				min.largestVowel = key
			} else {
				min.posStartingVowelRecur = key
				//min.largestVowel = key
			}
		}
	} else {
		if min.s >= 0 && min.largestConsonant < 0 {
			if min.posStartingVowelRecur >= 0 && s[min.s] != s[min.posStartingVowelRecur] {
				//clip
				clipMin(s, min, key)
			} else {
				min.e = key
				min.largestConsonant = key
			}
			updateMinBk(s, min)
		}
	}

}

func updateMinBk(s []rune, min *coord) {
	if isValidCoord(min) {
		if isValidCoord(MinBk) {
			tempBk := &coord{min.s, min.e, min.posStartingVowelRecur, min.largestConsonant, min.largestVowel}
			out := compare(s, MinBk, tempBk)
			if out > 1 {
				MinBk = tempBk
				//reset running min
				min = &coord{-1, -1, -1, -1, -1}
			}
		}
	}
}

func updateMaxBk(s []rune, max *coord) {
	if isValidCoord(max) {
		if isValidCoord(MaxBk) {
			tempBk := &coord{max.s, max.e, max.posStartingVowelRecur, max.largestConsonant, max.largestVowel}
			out := compare(s, MaxBk, tempBk)
			if out < 1 {
				MaxBk = tempBk
				//soft reset running max
				max.e = -1
			}
		}
	}
}

func adjustMax(s []rune, max *coord, key int) {
	if isVowel(s[key]) {
		if max.s < 0 {
			max.s = key
			max.largestVowel = key
		} else if s[max.s] < s[key] {
			updateMaxBk(s, max)
			//reset running max
			max.s = key
			max.e = -1
			max.largestVowel = key
		} else if s[max.s] == s[key] {
			max.posStartingVowelRecur = key
			max.largestVowel = key
		}
	} else {
		if max.s > -1 {
			if s[key] > s[max.largestVowel] {
				//clip
				clipMax(s, max, key)
			} else {
				max.e = key
			}
			updateMaxBk(s, max)
		}

	}
}

func clipMin(s []rune, c *coord, key int) {
	//clip
	c.s = c.posStartingVowelRecur
	c.e = key
	c.largestVowel = c.s
	c.posStartingVowelRecur = -1
	c.largestConsonant = key
}

func clipMax(s []rune, c *coord, key int) {
	//clip
	c.s = c.largestVowel
	c.e = key
	c.largestVowel = c.s
	c.posStartingVowelRecur = -1
	c.largestConsonant = key
}

func compare(s []rune, a, b *coord) int {
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

func isValidCoord(c *coord) bool {
	if c == nil {
		return false
	}
	return c.s > -1 && c.e > -1
}

func isVowel(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func closestPreviousVowel(c rune) rune {
	//alpha has a sorted vowel list
	var ret rune
	for k, _ := range alpha {
		if c > k {
			ret = k
		} else {
			return ret
		}
	}
	return ret
}

func getLocalMap(minVowel rune) map[rune]int {
	ret := make(map[rune]int)
	for k, _ := range alpha {
		if minVowel < k {
			ret[k] = 0
		}
	}
	return ret
}
