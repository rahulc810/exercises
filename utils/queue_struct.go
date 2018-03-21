package utils

import (
	"errors"
	"fmt"
	"sync"
)

type ObjQueue struct {
	arr  []interface{}
	head int
	tail int
	max  int
	size int
}

var mutex = &sync.Mutex{}

func NewObjQueue(max int) ObjQueue {
	return ObjQueue{arr: make([]interface{}, max), head: 0, tail: 0, max: max}
}

func (q *ObjQueue) Enqueue(item interface{}) error {
	mutex.Lock()
	defer mutex.Unlock()
	newTail := increment(q.tail, q.max)
	if q.tail == q.head {
		if q.arr[q.head] == nil {
			newTail = 0
		}
	}

	if newTail == q.head && q.arr[q.head] != nil {
		fmt.Printf("Queue is full failed to enqueue %d %d", q.head, newTail)
		return errors.New("Queue is full failed to enqueue")
	}
	q.tail = newTail
	s := item
	q.arr[q.tail] = s
	q.size++
	fmt.Printf("Queue =>%v\n", q)
	return nil
}

func (q *ObjQueue) Dequeue() (interface{}, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if q.head == q.tail && q.arr[q.head] == nil {
		//fmt.Printf("Cannot dequeue- its empty")
		return nil, errors.New("Cannot dequeue- its empty")
	}

	ret := q.arr[q.head]
	q.arr[q.head] = nil
	q.head = increment(q.head, q.max)
	q.size--
	return ret, nil
}

func (q *ObjQueue) Size() int {
	return q.size
}

func (q *ObjQueue) Max() int {
	return q.max
}

func increment(ht int, max int) int {
	if ht == max-1 {
		return 0
	}
	return ht + 1
}
