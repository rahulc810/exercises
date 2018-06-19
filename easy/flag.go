package easy

import (
	"fmt"
)

//dasd
func Flag() {
	s := "000011101100"
	c1 := 0
	c2 := 0
	total := 0
	curDir := "up"
	ignore := false
	r := []rune(s)[0]
	fmt.Println(len(s))

	for idx, ok := 1, true; ok; ok, idx = idx < len(s)-1, idx+1 {
		val := []rune(s)[idx]
		if curDir == "up" {
			if r == val {
				c1++
			} else {
				//beginning of down
				curDir = "down"
				c2 = 0
			}
		} else {
			if r == val {
				if c1 > 0 {
					c1--
				} else if !ignore {
					total += c2
					ignore = true
				}
				if c1 == 0 {
					total += c2
					ignore = true
				}

				c2++
			} else {
				//convert down to up
				total += c2
				curDir = "up"
				c1 = c2
				c2 = 1
				ignore = false
			}

		}
		fmt.Println(string(val), c1, c2, curDir, total)
		r = val
	}

	if !ignore {
		total += c2
	}

	fmt.Println(total)
}
