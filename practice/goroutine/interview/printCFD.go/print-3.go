package goroutine

import (
	"fmt"
	"sync"
)

// NOTE: 兩個協程交替列印10個字母和數字(結果: aaabbbb)
func PrintLetters() {
	letterCh := make(chan struct{}, 1)
	digitCh := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go printLetter(letterCh, digitCh, &wg)
		go printDigit(digitCh, letterCh, &wg)
	}

	//NOTE: 觸發第一個goroutine去執行
	letterCh <- struct{}{}

	wg.Wait()
	fmt.Println("執行結束")

}

func printLetter(ch <-chan struct{}, nextCh chan<- struct{}, wg *sync.WaitGroup) {
	<-ch
	fmt.Println("a")
	nextCh <- struct{}{}
	wg.Done()
}
func printDigit(ch <-chan struct{}, nextCh chan<- struct{}, wg *sync.WaitGroup) {
	<-ch
	fmt.Println("1")
	nextCh <- struct{}{}
	wg.Done()
}
