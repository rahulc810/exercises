package utils
import "fmt"
const(
	DefaultSize = 16
)

type List struct{
	arr []int
	size int
	cap int
}

func NewList(args...int) List{
	if len(args) == 0{
		return List{make([]int,DefaultSize),0,DefaultSize}
	}
		fmt.Printf("Creating new list og capacity: %d\n", args[0])
		return List{make([]int,args[0]),0,args[0]}	
}

func (l *List) Append(val int){
	fmt.Printf("Appending value: %d to list: %v\n", val, &l)
	if l.cap == l.size{
		fmt.Printf("List capacity full. Doubling the value\n")
		tempArr := make([]int,2*l.cap)
		for idx:=0;idx<len(l.arr);idx++{
			tempArr[idx] = l.arr[idx]
		}
		l.cap *= 2
		l.arr = tempArr
	}
	l.arr[l.size] = val
	l.size++ 
}