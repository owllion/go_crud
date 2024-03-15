package goroutine

import (
	"fmt"
	"time"
)

func Buffer2() {
	end := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Print goroutine")
		end <- true
	}()
	//NOTE: receive一定阻塞
	//詳情在buffer-1有紀錄了
	<-end
}
