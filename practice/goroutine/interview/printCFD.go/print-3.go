package goroutine

import (
	"fmt"
	"sync"
)

var letterCh = make(chan struct{})
var digitCh = make(chan struct{})
var wg = sync.WaitGroup{}

// NOTE: 兩個協程交替列印10個字母和數字(結果: aaabbbb)
func PrintLetters() {

	wg.Add(10)
	for i := 0; i < 5; i++ {
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
