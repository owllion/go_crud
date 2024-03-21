package goroutine

import (
	"fmt"
	"sync"
)

var letterCh2 = make(chan struct{}, 1)
var digitCh2 = make(chan struct{}, 1)
var wg2 = sync.WaitGroup{}

func PrintLetterAndDigit() {
	wg2.Add(2)
	//總共只創建兩條GR，各自用ch去通知對方可以print的時機，總共各10次
	///這方法比我的好，因為我那個創見了10條，也更難懂

	go printLetter2()
	go printDigit2()

	letterCh2 <- struct{}{}

	wg2.Wait()
	fmt.Println("執行結束")
}

func printLetter2() {
	for i := 0; i < 10; i++ {
		<-letterCh2
		fmt.Println("a")
		digitCh2 <- struct{}{}

	}
	defer wg2.Done() //這個defer一定要放在for迴圈外，不然到時候還沒for完，wg就會直接通知wg.Wait()可以放行，導致最後少print幾個內容!!!

}
func printDigit2() {
	for i := 0; i < 10; i++ {
		<-digitCh2
		fmt.Println("1")
		letterCh2 <- struct{}{}

	}
	defer wg2.Done() //這個defer一定要放在for迴圈外，不然到時候還沒for完，wg就會直接通知wg.Wait()可以放行，導致最後少print幾個內容!!!
}
