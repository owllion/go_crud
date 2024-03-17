package goroutine

import (
	"fmt"
)

var ch1 = make(chan struct{})
var ch2 = make(chan struct{})
var stop = make(chan struct{})
var count = 0

func Unbuffered1() {
	go func() {
		<-ch1
		fmt.Println("print ch1")
		ch2 <- struct{}{}
	}()

	go func() {
		<-ch2
		fmt.Println("print ch2")
		count += 1
		if count == 5 {
			stop <- struct{}{}
		}
	}()

	ch1 <- struct{}{}
	<-stop
	// end := make(chan bool)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Unbuffered-1 goroutine")
	// 	end <- true
	// }()

	// //NOTE: receive一定阻塞
	// <-end

}
