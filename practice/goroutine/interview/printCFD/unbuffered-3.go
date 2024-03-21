package goroutine

import (
	"fmt"
	"sync"
)

func Unbuffered3() {
	letter, number := make(chan bool), make(chan bool)

	wg := sync.WaitGroup{}
	go func() {

		for ch := 'A'; ch < 'Z'; ch += 2 {
			letter <- true
			fmt.Print(string(ch))
			fmt.Print(string(ch + 1))
			<-number

		}
		close(letter)
	}()

	wg.Add(1)
	go func() {
		start := 1
		for {
			number <- true

			fmt.Print(start)
			fmt.Print(start + 1)
			start += 2

			_, ok := <-letter
			if ok == false {
				break
			}
		}
		wg.Done()
	}()

	<-number

	wg.Wait()
	fmt.Print("\n")

}
