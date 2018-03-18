package easy

import(
	"fmt"
	"bufio"
	"strconv"
	"os"
)
//ALPHABET ...
var ALPHABET = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//BASE ...
var BASE = 62

func BaseMain(){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		input, _ := strconv.Atoi(scanner.Text())
		fmt.Printf("Output => %s \n", convert(input))
	}
}


func convert(input int) string{
	var ret string
	q:=input
	for ok:=true;ok;ok = (q!=0){
		r := q%BASE
		q=q/BASE
		ret += string(ALPHABET[r]) 
	}
	return ret
}