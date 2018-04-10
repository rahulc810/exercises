package utils

import (
	"fmt"
	"math"
	"reflect"
)

type Heap interface {
	push(key interface{})
	pop() (interface{}, bool)
	print()
}

type HeapArr struct {
	arr  []interface{}
	head int
	tail int
}

type Node struct {
	right *Node
	left  *Node
	key   interface{}
}

func (h *HeapArr) push(key interface{}) {
	fmt.Println("Executing push")
	h.arr = append(h.arr, key)
	child := h.tail
	h.maxHeapify(child)
	h.tail++

	h.printPretty()
}

func (h *HeapArr) pop() (interface{}, bool) {
	if h.tail == 0 {
		return nil, false
	}
	ret := h.arr[0]
	h.arr[0] = h.arr[h.tail-1]
	h.arr[h.tail-1] = "x"
	h.tail--
	h.Heapify()
	return ret, true
}

func (h *HeapArr) Heapify() {
	for i := (h.tail - 1) / 2; i >= 0; i-- {
		fmt.Println(i)
		h.maxHeapify(i)
	}
	h.printPretty()
}

func (h *HeapArr) maxHeapify(parent int) {
	var count uint32
	op := parent
	defer fmt.Printf("Max Heapify on [%v]: %v completed after %v hops\n", op, h.arr[op], count)
	for {
		leftChild := parent*2 + 1
		rightChild := parent*2 + 2

		rExists := rightChild < h.tail
		lExists := leftChild < h.tail

		var m int
		if rExists && lExists {
			m = max(h.arr, leftChild, rightChild)
		} else if lExists {
			m = leftChild
		} else if rExists {
			m = rightChild
		} else {
			return
		}
		if v, ok := compare(h.arr[parent], h.arr[m]); v < 1 && ok {
			swap(h.arr, parent, m)
			parent = m
		} else {
			return
		}
		count++
	}
}

func getParent(child int) int {
	if child%2 == 1 {
		return child / 2
	}
	return child/2 - 1
}

func (h *HeapArr) printPretty() {
	filler := " "
	l := len(h.arr)
	lvls := int(math.Ceil(math.Log2(float64(l + 1))))
	nodesOnLastLvl := 1 << (uint64(lvls - 1))
	//tPossibleNodes := (1 << (uint32(lvls))) - 1
	//fmt.Println("Total levels: ", lvls)
	//fmt.Println("Nodes in last level: ", nodesOnLastLvl)
	//fmt.Println("Total possible nodes in the tree: ", tPossibleNodes)
	tCols := (nodesOnLastLvl-1)*3 + nodesOnLastLvl //spacesBetweenNodes for last lvl is alway 3
	idx := 0
	for lv := 1; lv <= lvls; lv++ {
		if lv == 1 {
			//first level is special
			p(filler, (tCols-1)/2)
			p(h.arr[idx], 1)
			idx++
			p(filler, (tCols-1)/2)
			p("\n", 1)
			continue
		}
		//prepare arrows towards current level (displayed above the current level)
		numberOfNodes := 1 << (uint32(lv) - 1)
		spacesBetweenNodes := (1 << (uint32(lvls - lv + 2))) - 1
		spacesOnEnds := tCols - ((numberOfNodes - 1) * spacesBetweenNodes) - numberOfNodes
		nodesOnPreviousLvl := 1 << uint32(lv-2)

		/*for each row :
		print spacesOnEnd/2
			for each node in the previous level
				print 1 + start padding + / + middle padding + \ + end padding + 1
				followed by spaces between nodes of current lvl except for the last iteration
		print spacesOnEnd/2
		*/
		for ro := 1; ro <= (1<<uint32(lvls-lv+1))-1; ro++ {
			p(filler, spacesOnEnds/2)
			for nodes := 1; nodes <= nodesOnPreviousLvl; nodes++ {
				p(filler, 1)
				p(filler, (spacesBetweenNodes/2)-ro)
				p("/", 1)
				p(filler, 2*ro-1)
				p("\\", 1)
				p(filler, (spacesBetweenNodes/2)-ro)
				p(filler, 1)
				if nodes < nodesOnPreviousLvl {
					p(filler, spacesBetweenNodes)
				}
			}
			p(filler, spacesOnEnds/2)
			p("\n", 1)
		}

		//print node row
		p(filler, spacesOnEnds/2)
		for s := 0; s < numberOfNodes; s++ {
			if idx >= l {
				p("x", 1)
			} else {
				p(h.arr[idx], 1)
			}
			idx++
			if s < numberOfNodes-1 {
				p(filler, spacesBetweenNodes)
			}
		}
		p(filler, spacesOnEnds/2)
		p("\n", 1)
	}
}

func (h *HeapArr) print() {
	filler := "*"
	l := len(h.arr)
	lvls := int(math.Ceil(math.Log2(float64(l))))
	nodesOnLvl := 1 << (uint64(lvls))
	fmt.Println("Total levels: ", lvls)
	fmt.Println("Elements in last level: ", nodesOnLvl)

	idx := 0
	for lv := 1; lv <= lvls; lv++ {
		if lv == 1 {
			ws := (nodesOnLvl - 1) / 2
			for s := 0; s < ws; s++ {
				fmt.Printf(filler)
			}
			fmt.Printf("%v", h.arr[idx])
			idx++
			for s := 0; s < ws; s++ {
				fmt.Printf(filler)
			}
		} else {
			numberOfNodes := 1 << (uint32(lv) - 1)
			spacesBetweenNodes := (1 << (uint32(lvls - lv + 1))) - 1
			spacesOnEnds := nodesOnLvl - ((numberOfNodes - 1) * spacesBetweenNodes) - numberOfNodes
			for s := 0; s < spacesOnEnds/2; s++ {
				fmt.Printf(filler)
			}
			for s := 0; s < numberOfNodes && idx < l; s++ {
				fmt.Printf("%v", h.arr[idx])
				idx++
				for w := 0; w < spacesBetweenNodes && s < numberOfNodes-1; w++ {
					fmt.Printf(filler)
				}
			}
			for s := 0; s < spacesOnEnds/2; s++ {
				fmt.Printf(filler)
			}

		}
		fmt.Println()
	}

}

func p(filler interface{}, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%v", filler)
	}
}

func max(arr []interface{}, a, b int) int {
	val := arr[a]
	switch val.(type) {
	case int:
		aInt, _ := arr[a].(int)
		bInt, _ := arr[b].(int)

		if aInt > bInt {
			return a
		}
		return b
	default:
		return 0
	}
}

func compare(a, b interface{}) (int, bool) {
	aType := reflect.TypeOf(a)
	bType := reflect.TypeOf(b)
	if aType != bType {
		return 0, false
	}
	switch a.(type) {
	case int:
		aInt, oka := a.(int)
		bInt, okb := b.(int)

		if !oka || !okb {
			return 0, false
		}
		if aInt > bInt {
			return 1, true
		} else if aInt < bInt {
			return -1, true
		} else {
			return 0, true
		}
	default:
		return 0, false
	}

}

func swap(arr []interface{}, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func NewHeap() Heap {
	arr := make([]interface{}, 0)
	return &HeapArr{arr, 0, 0}
}

func Exec() {
	sl := []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q"}
	sl = []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
	sl = []interface{}{6, 5, 3, 1, 8, 7, 2, 4}
	h := &HeapArr{sl, 0, 8}

	//h.printPretty()
	h.Heapify()

	p, _ := h.pop()
	fmt.Printf("Popped: %v\n", p)

	p, _ = h.pop()
	fmt.Printf("Popped: %v\n", p)

	//	h.push(18)
}
