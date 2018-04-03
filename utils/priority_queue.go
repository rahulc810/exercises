package utils

import (
	"fmt"
	"math"
)

type Heap interface {
	push(key interface{})
	pop() (interface{}, bool)
	print()
}

type HeapArr struct {
	arr  []interface{}
	head int
}

type Node struct {
	right *Node
	left  *Node
	key   interface{}
}

func (h *HeapArr) push(key interface{}) {
	fmt.Println("Executing push")
}

func (h *HeapArr) pop() (interface{}, bool) {
	return nil, false
}

func preToIn(pre []interface{}, in []interface{}) {
	//idx := 0
	for i := 0; i < len(pre); i++ {
		//
	}
}
func (h *HeapArr) printPretty() {
	filler := " "
	l := len(h.arr)
	lvls := int(math.Ceil(math.Log2(float64(l))))
	nodesOnLastLvl := 1 << (uint64(lvls - 1))
	tPossibleNodes := (1 << (uint32(lvls))) - 1
	fmt.Println("Total levels: ", lvls)
	fmt.Println("Nodes in last level: ", nodesOnLastLvl)
	fmt.Println("Total possible nodes in the tree: ", tPossibleNodes)
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

func NewHeap() Heap {
	arr := make([]interface{}, 0)
	return &HeapArr{arr, 0}
}

func Exec() {
	sl := []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q"}
	sl = []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
	h := &HeapArr{sl, 0}
	h.printPretty()

}
