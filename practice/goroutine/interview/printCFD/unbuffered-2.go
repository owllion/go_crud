package goroutine

import (
	"fmt"
	"time"
)

func Unbuffered2() {
	end := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Unbuffered-2 goroutine")
		<-end //goruotine會等到這邊收到值才結束，最終結果和Unbuffered1一樣喔
		//但你可能會想問，為啥main goruotine不會直接結束? main哪會管這個gr執行完畢了沒?
		//會等是由於unbuffered的特性，因為它是讀寫blocking(其中一方還沒準備好，該行就會暫停執行)，所以 end <- true執行到的時候，要是end is not ready for execution(still sleeping)，那就會卡在 end <- true這一行，不會繼續執行喔!
		//最終等2秒，印出字串，end才準備好接收，最後才執行完畢!
	}()

	end <- true
}
