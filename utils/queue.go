package utils

import(
	"errors"
	"fmt"
)

type StringQueue struct{
 arr []string
 head int
 tail int
 max int
}

type Queue interface{
	Enqueue(item interface{}) error
	Dequeue() (interface{},error)
	Size() int
	Max() int
}

func NewQueue(max int) StringQueue{
	return StringQueue{arr : make([]string, max), head: 0, tail: 0, max : max}
}

func (q *StringQueue) Enqueue(item interface{}) error{
	if q.tail == q.head && q.arr[q.head]==""{
		s,ok := item.(string)
		if !ok{
			return errors.New("Bad item to enqueue")	
		}
		q.arr[q.tail] = s
		return nil
	}


	newTail := incrementHeadOrTail(q.tail,q.Size())
	if newTail == q.head{
		fmt.Printf("Queue is full failed to enqueue %d %d", q.head, newTail)
		return errors.New("Queue is full failed to enqueue")
	}
	q.tail = newTail
	s,ok := item.(string)
	if !ok{
		return errors.New("Bad item to enqueue")	
	}
	q.arr[q.tail] = s
	return nil
}

func (q *StringQueue) Dequeue() (interface{},error){
	if q.head==q.tail && q.arr[q.head]== ""{
		fmt.Printf("Cannot dequeue- its empty")
		return nil, errors.New("Cannot dequeue- its empty")
	}

	ret := q.arr[q.head]
	q.arr[q.head] = ""
	q.head = incrementHeadOrTail(q.head, q.Size())
	return ret,nil
}

func incrementHeadOrTail(ht int, size int) int{
	fmt.Printf("Size %d\n", size)
	 if ht == size -1{
		 return 0
	 }
		 return ht+1
}

func (q *StringQueue) Size() int{
	if q.arr[q.head]==q.arr[q.tail]{
		if q.arr[q.head] == ""{
			return 0
		}
		return 1
	}

	if q.head > q.tail{
		return len(q.arr) -q.head + q.tail + 1
	}

	return q.tail - q.head + 1 
}

func (q *StringQueue) Max() int{
	return q.max
}