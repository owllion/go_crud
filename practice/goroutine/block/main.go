package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// 启动一个 goroutine 来处理接收操作
	go func() {
		fmt.Println("in first goruotine------------")
		for {
			select {
			case x := <-ch:
				fmt.Println("Received value:", x)
			default:
				fmt.Println("No value received")
			}
			//NOTE: 如果不加上這個，也就是部稍微等等，那傳送方還沒準備好，這個 x:= <- ch就會永遠阻塞，永遠不會完成，導致只會印出No value received.
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// 在没有足够时间准备接收方的情况下，尝试发送值到通道
	go func() {
		fmt.Println("in sec goruotine----")
		time.Sleep(1 * time.Second) // 等待一秒钟，模拟没有足够的时间给接收方准备
		//NOTE:不加上迴圈的話，代表只會嘗試傳輸值依次，那要是接收方尚未準備好，依樣永遠會阻塞，永遠不可能傳出去
		for {
			// time.Sleep(500 * time.Millisecond) // 等待一段时间，确保接收操作已经开始
			select {
			case ch <- 42:
				fmt.Println("Value 42 sent to channel")
			default:
				fmt.Println("No value sent")
			}
			//NOTE: 和接收方一樣的原因，不寫這，到時候接收方還沒準備好(準備好接收)，又因這次unbuffered，讀和寫必須是一組的(有讀就必須要有寫，反之)，所以有一方沒準備好，當然 ch<-42會永遠阻塞，最後只會印出No value sent!!!!!
			time.Sleep(100 * time.Millisecond)
		}

	}()

	//NOTE: 稅2秒起來後，就會繼續執行main goroutine囉!
	time.Sleep(2 * time.Second) // 等待足够长的时间，确保能够看到输出结果
	fmt.Println("執行結束")
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// 阻塞寫入操作
// 	ch := make(chan int)

// 	// 非阻塞讀取操作
// 	go func() {
// 		for {
// 			select {
// 			case x := <-ch:
// 				fmt.Println("Received value:", x)
// 			default:
// 				fmt.Println("No value received")
// 			}
// 			// time.Sleep(100 * time.Millisecond) // 降低 CPU 使用率
// 		}
// 	}()

// 	time.Sleep(1 * time.Second) // 讓 goroutine 有機會開始執行
// 	fmt.Println("睡眠一秒後")
// 	// 阻塞寫入操作
// 	ch <- 42
// 	fmt.Println("Value sent to channel")

// 	// 非阻塞寫入操作
// 	select {
// 	case ch <- 10: //不會跑這個case是因為ch是unbuffered的，所以因為這邊沒有和他對應的讀取操作，他就會block直到有對應的接收方，因此永遠不會是"完成"狀態
// 		fmt.Println("Value 10 sent to channel")
// 	default:
// 		fmt.Println("No value sent")
// 	}
// 	fmt.Println("執行結束")
// }
