package main

import (
	"fmt"
	"time"
)

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("RECOVER:", r) //印出的就會是傳入panic的內容(即WOWO)
	}

}

func main() {
	//NOTE:在更強一點提到，recover一定要搭配defer使用才能順利拿回panic後的控制權，這是因為panic的期間，他只會執行defer相關函數(不管怎樣都會幫你去把該關閉的東西都給關了)，如果有defer以外的都不會去執行，這樣懂了吧~所以不佳defer就根本不會被讀到啦
	go func() {
		defer handlePanic()
		fmt.Println("gr1")
	}()
	go func() {
		defer handlePanic()
		fmt.Println("gr2")
		panic("WOWO")
	}()
	go func() {
		defer handlePanic()
		fmt.Println("gr3")
	}()
	go func() {
		defer handlePanic()
		fmt.Println("gr4")
	}()
	go func() {
		defer handlePanic()
		fmt.Println("gr5")
	}()
	go func() {
		defer handlePanic()
		fmt.Println("gr6")
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("recover from panic!!")
}
