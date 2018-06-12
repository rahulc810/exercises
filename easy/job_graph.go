package easy

import (
	"fmt"
	"sync"
	"time"
)

var idle int
var wg sync.WaitGroup

type Node struct {
	name     string
	parent   []string
	children []string
	weight   time.Duration
	status   bool
}

var lk map[string]Node

//
func JobGraph() {
	fmt.Println("Prepare graph")
	a := Node{"A", nil, nil, 2, false}
	b := Node{"B", nil, nil, 2, false}
	c := Node{"C", nil, nil, 2, false}
	d := Node{"D", nil, nil, 2, false}
	e := Node{"E", nil, nil, 2, false}
	f := Node{"F", nil, nil, 2, false}
	g := Node{"G", nil, nil, 2, false}
	h := Node{"H", nil, nil, 2, false}

	a.children = []string{"B"}
	b.children = []string{"C", "D"}
	b.parent = []string{"A"}
	c.children = []string{"E", "F"}
	c.parent = []string{"B"}
	d.children = []string{"H"}
	d.parent = []string{"B"}
	e.children = []string{"G"}
	e.parent = []string{"C"}
	f.children = []string{"G"}
	f.parent = []string{"C"}
	g.children = []string{"H"}
	g.parent = []string{"F", "E"}
	h.parent = []string{"D", "G"}

	lk = make(map[string]Node, 100)

	lk["A"] = a
	lk["B"] = b
	lk["C"] = c
	lk["D"] = d
	lk["E"] = e
	lk["F"] = f
	lk["G"] = g
	lk["H"] = h

	//graph prepared
	eval(a)
}

func eval(head Node) {
	jobs := make(chan Node, 100)
	wg.Add(8)

	idle = 1
	CPU(idle, jobs)

	jobs <- head
	start := time.Now()
	wg.Wait()
	end := time.Now()

	fmt.Printf("Time taken - %d", end.Sub(start)/1000000000)

}

func CPU(cores int, jobs chan Node) {
	for i := 0; i < cores; i++ {
		go func() {
			for j := range jobs {
				time.Sleep(j.weight * time.Second)
				j.status = true
				lk[j.name] = j
				fmt.Printf("%v complete \n", j.name)
				//Queuechildren once parent is finished
				for _, child := range j.children {
					enqueue := true
					for _, parent := range lk[child].parent {
						if !lk[parent].status {
							enqueue = false
							break
						}
					}
					if enqueue {
						//q.Enqueue(child)
						fmt.Printf("%v added \n", child)
						jobs <- lk[child]
					}
				}
				wg.Done()
			}
		}()
	}
}
