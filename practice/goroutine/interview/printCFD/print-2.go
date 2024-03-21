package goroutine

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

//NOTE:題目
//有三個函數，分別列印"cat", "fish","dog"要求每個函數都用一個goroutine，每一個都列印100次，例如cat*100 → fish*100 …etc

var Wg = &sync.WaitGroup{}

func PrintTotal100Times() {
	// catCh := make(chan struct{}, 1)
	// fishCh := make(chan struct{}, 1)
	// dogCh := make(chan struct{}, 1)
	catCh := make(chan struct{}, 1)
	fishCh := make(chan struct{}, 1)
	dogCh := make(chan struct{}, 1)

	Wg.Add(15)

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
	for i := 0; i < 5; i++ {
		go printCat1Times(catCh, fishCh)
		go printFish1Times(fishCh, dogCh)
		go printDog1Times(dogCh, catCh)
	}

	catCh <- struct{}{}

	//NOTE:補充
	//用一個nuffered channel去紀錄當前已經跑幾次了，滿30即可讓主程式繼續執行
	// for i := 0; i < 10; i++ {
	// 	<-done
	// }
	Wg.Wait()
	fmt.Println("執行結束")
	//NOTE: 不能用這種寫法，因為執行到的10次的時候，其實那只有第一個cat有被執行，for迴圈就跳出了，然後印出執行結束，但是另外兩個gr其實也還在執行，所以有機率會導致"執行結束"後還印出dog、fish的情況。應該還是要用waitGroup
}

func printCat1Times(ch <-chan struct{}, fishCh chan<- struct{}) {
	<-ch
	fmt.Println("cat")
	fishCh <- struct{}{}
	defer Wg.Done()
}
func printFish1Times(ch <-chan struct{}, dogCh chan<- struct{}) {
	<-ch
	fmt.Println("fish")
	dogCh <- struct{}{}
	defer Wg.Done()
}
func printDog1Times(ch <-chan struct{}, catCh chan<- struct{}) {
	<-ch
	fmt.Println("dog")
	catCh <- struct{}{}
	defer Wg.Done()
}
