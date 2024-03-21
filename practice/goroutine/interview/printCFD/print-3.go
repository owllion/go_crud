package goroutine

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

var letterCh = make(chan struct{}, 1)
var digitCh = make(chan struct{}, 1)
var wg = sync.WaitGroup{}

// NOTE: 兩個協程交替列印10個字母和數字(結果: aaabbbb)
func PrintLetters() {
	file, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = trace.Start(file)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	wg.Add(10)
	for i := 0; i < 5; i++ { //NOTE: 這回圈會先把所有gr一次創建完，然後就會照著創建的順序去執行、exit
		go printLetter()
		go printDigit()
	}

	//NOTE: 觸發第一個goroutine去執行
	letterCh <- struct{}{}

	wg.Wait()
	fmt.Println("執行結束")

}

func printLetter() {
	<-letterCh
	fmt.Println("a")
	digitCh <- struct{}{}

	// defer func() {
	// 	wg.Done()
	// 	fmt.Println("letter goruotine done")
	// }()
	defer wg.Done()

}
func printDigit() {
	<-digitCh
	fmt.Println("1")

	letterCh <- struct{}{}

	// defer func() {
	// 	wg.Done()
	// 	fmt.Println("digit goruotine done")
	// }()
	defer wg.Done()

}

//TODO: 測試傳入值到unbuffered，然後直接關閉的結果
