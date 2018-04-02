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

func NewHeap() Heap {
	arr := make([]interface{}, 0)
	return &HeapArr{arr, 0}
}

func Exec() {
	sl := []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	h := &HeapArr{sl, 0}
	h.print()

}
