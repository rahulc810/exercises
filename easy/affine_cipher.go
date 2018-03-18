package easy

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"regexp"
	"sync"
)

var ASCIIALPHABET string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func AffixCipherMain(){
	pipe := make(chan int,100)
	var dict []string
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		dict = loadAsDict("/Users/rahchauh/Desktop/enable1.txt", pipe)
		wg.Done()
	}()
		
	wg.Wait()
	close(pipe)
	scnr := bufio.NewScanner(os.Stdin)
	for scnr.Scan(){
		token:=scnr.Text()
		if len(token)==0{
			break
		}
		var dc string
		var toBeRanked []string
		var maxRank int = 0
		var secondMaxRank int = 0
		var maxText string = ""
		var secondMaxText string = ""
		for _,a := range []int{3,5,7,11,15,17,19,21,23,25}{
			for b:=1;b< 27;b++{
				dc = decrypt(a,b,token)
				fmt.Printf("%s\n",dc)
				toBeRanked = append(toBeRanked, dc)
			}
		}

		fmt.Printf("\nDictionary created - pipe closed\n")
		for _,dcText := range(toBeRanked){
			points:=rank(&dict, dcText)
			if points > maxRank{
				maxRank = points
				maxText = dcText
			}else if points > secondMaxRank{
				secondMaxRank = points
				secondMaxText = dcText
			}
		}
		fmt.Printf("[%d]%s", maxRank, maxText)
		fmt.Printf("[%d]%s", secondMaxRank, secondMaxText)
	}

}

func decrypt(a int, b int, text string) string{
var out string
for _,val:= range text{
	if val == ','|| val == '/' || val ==' '{
		out += string(val)
	}else{
		x:= int(val)-'A'
		mul:= a*x + b
		adj := mod(mul,26)
		
		out += string(adj+'A')
	}
}
return out
}
func mod(z,m int) int{
	if z >-1 {
		return z%m
	}
z = -z
fmt.Printf("DEBUG string value of updated %d \n", z)
return (z/m + 1)*m - z

}

func tokenize(text string)[]string{
	return regexp.MustCompile("[,/ ]").Split(text,-1)
}

func stringInSlice(arr *[]string, val string) bool{
	end := len(*arr)-1
	start:=0 
	index := end/2 +1;

	for (index <= end && index >= start){
		if (*arr)[index]==val{
			return true
		}else if end-start<2{
			return (*arr)[end] ==val || (*arr)[start]==val
		}else if (*arr)[index] > val{
			end = index
		}else if (*arr)[index] < val{
			start = index
		}
		index = (end-start)/2 + start
	}
	return false
}

func loadAsDict(location string, pipe chan int) []string{
	fmt.Printf("LOADING...")
	file, err := os.Open(location)
	defer file.Close()
	if err != nil{
		fmt.Printf("Unable to load file into dict : %s", location)
		os.Exit(1)
	}
	dict := make([]string, 172820)
	scnr := bufio.NewScanner(file)
	var token string
	var i int
	for scnr.Scan(){
		token = scnr.Text()
		dict[i] = token
		i++
	}
	pipe <- 100
	return dict
}

func rank(dict *[]string, dcText string) int{
	points := 1
	tokens := tokenize(dcText)
	for _,token := range tokens{
		isPresent := stringInSlice(dict, strings.ToLower(token))
		if isPresent{
			points++
		}
	}
	return points
}