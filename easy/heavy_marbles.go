package easy

import (
	"fmt"
	"math/rand"
	"time"
)

type tuple struct {
	arr        []int
	start, end int
}

func (t *tuple) weigh() int {
	var w int
	for i := t.start; i <= t.end; i++ {
		w += t.arr[i]
	}
	return w
}

func (t *tuple) length() int {
	return t.end - t.start + 1
}

func solveMarbles(t tuple) (int, tuple) {
	n := t.length()
	switch n {
	case 1:
		return 0, t
	case 2:
		return 1, t
	case 3:
		return 1, t
	default:
		t1 := tuple{t.arr, 0 + t.start, n/3 - 1 + t.start}
		t2 := tuple{t.arr, n/3 + t.start, 2*n/3 - 1 + t.start}
		t3 := tuple{t.arr, 2*n/3 + t.start, t.end}
		var target tuple
		if t1.weigh() > t2.weigh() {
			target = t1
		} else if t1.weigh() < t2.weigh() {
			target = t2
		} else {
			target = t3
		}
		count, t := solveMarbles(target)
		count++
		fmt.Printf("Retuned =>%v\n ", t)
		return count, t
	}
}

func MarblePuzzle(n int) {
	in := make([]int, n)

	rand.Seed(time.Now().Unix())
	idx := rand.Intn(n)
	in[idx] = 1

	fmt.Println(solveMarbles(tuple{in, 0, n - 1}))

}
