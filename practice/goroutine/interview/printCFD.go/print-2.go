package goroutine

import (
	"fmt"
)

//NOTE:題目
//有三個函數，分別列印"cat", "fish","dog"要求每個函數都用一個goroutine，每一個都列印100次，例如cat*100 → fish*100 …etc

func PrintTotal100Times() {
	catCh := make(chan struct{})
	fishCh := make(chan struct{})
	dogCh := make(chan struct{})
	done := make(chan struct{}, 10)

	for i := 0; i < 4; i++ {
		go printCat1Times(catCh, fishCh, done)
		go printFish1Times(fishCh, dogCh, done)
		go printDog1Times(dogCh, catCh, done)
	}

	catCh <- struct{}{}

	//NOTE:補充
	//用一個nuffered channel去紀錄當前已經跑幾次了，滿30即可讓主程式繼續執行
	for i := 0; i < 10; i++ {
		<-done
	}
	fmt.Println("執行結束")
}

func printCat1Times(ch <-chan struct{}, fishCh chan<- struct{}, done chan<- struct{}) {
	<-ch
	fmt.Println("cat")
	fishCh <- struct{}{}
	done <- struct{}{}
}
func printFish1Times(ch <-chan struct{}, dogCh chan<- struct{}, done chan<- struct{}) {
	<-ch
	fmt.Println("fish")
	dogCh <- struct{}{}
	done <- struct{}{}

}
func printDog1Times(ch <-chan struct{}, catCh chan<- struct{}, done chan<- struct{}) {
	<-ch
	fmt.Println("dog")
	catCh <- struct{}{}
	done <- struct{}{}

}
