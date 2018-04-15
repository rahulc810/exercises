package utils

import (
	"fmt"
	"math"
)

const (
	LEFT  = 0
	RIGHT = 1
)

func lr(a *Node) *Node {
	if a == nil || a.right == nil || a.right.right == nil {
		return nil
	}
	b := a.right
	b.left = a
	a.right = nil
	return b
}

func rr(a *Node) *Node {
	if a == nil || a.left == nil || a.left.left == nil {
		return nil
	}
	b := a.left
	a.left = nil
	b.right = a
	return b
}

func dr(a *Node) *Node {
	if a == nil || a.left == nil || a.left.right == nil {
		return nil
	}
	b, c := a.left, a.left.right
	a.left, b.right = nil, nil
	c.left, c.right = b, a
	return c
}

func dl(a *Node) *Node {
	if a == nil || a.right == nil || a.right.left == nil {
		return nil
	}
	b, c := a.right, a.right.left
	a.right, b.left = nil, nil
	c.left, c.right = b, a
	return c
}

func Insert(root, c *Node) *Node {

	if root == nil {
		return c
	}

	printTree(root)
	fmt.Println("Now inserting : ", c.key)
	var su, a, b *Node
	dirF, dirS := LEFT, LEFT

	for b = root; ; {
		if Compare(b.key, c.key) < 1 {
			if b.right != nil {
				su = a
				a = b
				b = b.right
				dirF = dirS
				dirS = RIGHT
			} else {
				b.right = c
				r := rotate(a, dirS, RIGHT)
				if dirF == LEFT && su != nil {
					su.left = r
				} else if su != nil {
					su.right = r
				} else if a != nil {
					return r
				}
				return root
			}
		} else {
			if b.left != nil {
				su = a
				a = b
				b = b.left
				dirF = dirS
				dirS = LEFT
			} else {
				b.left = c
				r := rotate(a, dirS, LEFT)
				if dirF == LEFT && su != nil {
					su.left = r
				} else if su != nil {
					su.right = r
				} else if a != nil {
					return r
				}
				return root
			}
		}
	}
	return root
}

func rotate(a *Node, x, y int) *Node {
	fmt.Println("*****")
	printTree(a)
	var ret *Node
	fmt.Println("----")
	defer printTree(ret)
	if x == 1 {
		if y == 1 {
			ret = lr(a)
		} else {
			ret = dl(a)
		}
	} else {
		if y == 1 {
			ret = dr(a)
		} else {
			ret = rr(a)
		}
	}
	return ret
}

func printTree(root *Node) {

	if root == nil {
		return
	}

	read := make([]interface{}, 0)
	use := make(chan *Node, 1000)
	cur := root
	read = append(read, cur)
	for ; cur != nil; cur = <-use {
		if cur.left == nil {
			read = append(read, &Node{nil, nil, 0})
		} else {
			read = append(read, cur.left)
			use <- cur.left
		}
		if cur.right == nil {
			read = append(read, &Node{nil, nil, 0})
		} else {
			read = append(read, cur.right)
			use <- cur.right
		}

		if len(use) == 0 {
			break
		}
	}

	filler := " "
	l := len(read)
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
			p(read[idx].(*Node).key, 1)
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
				p(read[idx].(*Node).key, 1)
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
