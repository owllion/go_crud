package goroutine

import "fmt"

//NOTE:題目
//有三個函數，分別列印"cat", "fish","dog"要求每個函數都用一個goroutine，每一個都列印100次，例如cat*100 → fish*100 …etc

func PrintStr() {
	catCh := make(chan struct{})
	fishCh := make(chan struct{})
	dogCh := make(chan struct{})
	stop := make(chan struct{})
	go printCat(catCh, fishCh)
	go printFish(fishCh, dogCh)
	go printDog(dogCh, stop)
	catCh <- struct{}{}

	<-stop
	fmt.Println("執行結束")
}

func printCat(ch <-chan struct{}, fishCh chan<- struct{}) {
	<-ch
	for i := 0; i < 5; i++ {
		fmt.Println("cat")
	}
	fishCh <- struct{}{}

}
func printFish(ch <-chan struct{}, dogCh chan<- struct{}) {
	<-ch
	for i := 0; i < 5; i++ {
		fmt.Println("fish")
	}
	dogCh <- struct{}{}

}
func printDog(ch <-chan struct{}, stop chan<- struct{}) {
	<-ch
	for i := 0; i < 5; i++ {
		fmt.Println("dog")
	}
	stop <- struct{}{}

}
