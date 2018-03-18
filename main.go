package main

import (
	"fmt"
	"strconv"
	"strings"
	//"github.com/rahulc810/exercises/utils"
	"github.com/rahulc810/exercises/easy"
)

func main() {
easy.Listen()
//clockMain()
//easy.BaseMain()
//easy.AffixCipherMain()
//easy.RabbitMain()
}

func clockMain(){
inputs := []string{"00:00", "03:45", "07:01", "13:19"}
for idx, val := range inputs {
	s := strings.Split(val, ":")
	hh, _ := strconv.Atoi(s[0])
	mm, _ := strconv.Atoi(s[1])

	fmt.Printf("[%d] Parsed input<= %d:%d\n", idx, hh, mm)

	hours, phase := easy.ConvertHours(hh)

	fmt.Printf("[%d] Output=> %s %s %s \n", idx, easy.TranslateToWords(hours, false), easy.TranslateToWords(mm, true), phase)

}
}
