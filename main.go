package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	//"github.com/rahulc810/exercises/utils"

	"github.com/rahulc810/exercises/easy"
)

func main() {
	//defer profile.Start(profile.MemProfile).Stop()
	easy.FindDctSubString()
	//utils.Exec()
	//easy.MarblePuzzle(89)
	//hikeServer()
	//easy.Listen() -- hike server listen
	//clockMain()
	//easy.BaseMain()
	//easy.AffixCipherMain()
	//easy.RabbitMain()
}

func hikeServer() {
	easy.Start()
	id1 := easy.Put([]byte("ABC"))
	id2 := easy.Put([]byte("ABC"))
	id3 := easy.Put([]byte("ABC"))
	id4 := easy.Put([]byte("ABC"))

	for i := 0; i < 1000000; i++ {
	}
	runtime.GC()
	easy.Del(id1)
	easy.Del(id2)
	easy.Del(id3)
	easy.Del(id4)
	for {
	}
}

func clockMain() {
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
