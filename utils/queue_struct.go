package utils

import(
	"errors"
	"fmt"
	"sync"
)

type ObjQueue struct{
 arr []interface{}
 head int
 tail int
 max int
}

var mutex = &sync.Mutex{}

func NewObjQueue(max int) ObjQueue{
	return ObjQueue{arr : make([]interface{}, max), head: 0, tail: 0, max : max}
}

func (q *ObjQueue) Enqueue(item interface{}) error{
	mutex.Lock()
	defer mutex.Unlock()
	if q.tail == q.head && q.arr[q.head]==nil{
		s:= item
		q.arr[q.tail] = s
		return nil
	}


	newTail := incrementHeadOrTail(q.tail,q.Size())
	if newTail == q.head{
		fmt.Printf("Queue is full failed to enqueue %d %d", q.head, newTail)
		return errors.New("Queue is full failed to enqueue")
	}
	q.tail = newTail
	s:= item
	q.arr[q.tail] = s
	return nil
}

func (q *ObjQueue) Dequeue() (interface{},error){
	mutex.Lock()
	defer mutex.Unlock()
	if q.head==q.tail && q.arr[q.head]== nil{
		//fmt.Printf("Cannot dequeue- its empty")
		return nil, errors.New("Cannot dequeue- its empty")
	}

	ret := q.arr[q.head]
	q.arr[q.head] = nil
	q.head = incrementHeadOrTail(q.head, q.Size())
	return ret,nil
}

func (q *ObjQueue) Size() int{
	if q.arr[q.head]==q.arr[q.tail]{
		if q.arr[q.head] == nil{
			return 0
		}
		return 1
	}

	if q.head > q.tail{
		return len(q.arr) -q.head + q.tail + 1
	}

	return q.tail - q.head + 1 
}

func (q *ObjQueue) Max() int{
	return q.max
}