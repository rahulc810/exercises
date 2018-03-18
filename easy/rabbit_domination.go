package easy

import (
	"fmt"
	"strconv"
	"github.com/rahulc810/exercises/utils"
)

func RabbitMain(){
	q:= utils.NewQueue(98) 
	qm:=utils.NewQueue(98)  
	for i:=0;i<93;i++{
		q.Enqueue("0")
		qm.Enqueue("0")
	}
	//First rabbit
	q.Enqueue("4")
	q.Enqueue("0")
		q.Enqueue("0")

			qm.Enqueue("0")
				qm.Enqueue("0")
					qm.Enqueue("0")

	result := 1000000000
	total := 4
	f2 := 4
	i:=0
	var fa3,f0,f1,f3,fDie,fBirth int
	var mBirth,mDie int

	for total < result{
		v,_ := q.Dequeue()
		fDie, _ = strconv.Atoi(string(v.(string)))
		fmt.Printf("\nFdie : %d", fDie)
		//fmt.Printf("\nFemale channel length : %d", len(q))
		v,_ = qm.Dequeue()
		mDie,_ = strconv.Atoi(string(v.(string)))
		fBirth = fa3*9
		mBirth = fa3*5
		q.Enqueue(strconv.Itoa(fBirth))
		qm.Enqueue(strconv.Itoa(mBirth))
		fa3 += f3-fDie
		f3=f2
		f2=f1
		f1=f0
		f0=fBirth
		total += mBirth -mDie + fBirth -fDie
		i++
	}

	fmt.Printf("Total ===> %d \n",total)
	fmt.Printf("Iterations ===> %d \n",i)
}



