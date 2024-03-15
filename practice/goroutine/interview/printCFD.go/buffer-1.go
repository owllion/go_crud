package goroutine

import (
	"fmt"
	"time"
)

func Buffer1() {
	end := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Goroutine.")
		<-end //receive一定阻塞
	}()
	//但這是buffered ch,只有讀取(為空)或是ch滿了才會阻塞，也就是說
	//你在CH滿之前放入都不會阻塞，也不用特地去取出
	//所以這個Buffer1函數，開完goroutine，執行到這一行，它就會直接結束了
	//因為它不會阻塞阿!
	//但是Buffer2則是把這兩位置互換， <-end換到這一行，剛說了讀取(receive)一定阻塞，所以等了兩秒鐘，end收到值之後才結束執行，print Goroutine也成功被執行到!
	end <- true
}
