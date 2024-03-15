package main

import "fmt"

//觀察defer執行順序

func namedReturn() (num int) {
	num = 0
	for num < 20 {
		num++
	}
	//NOTE: defer定義&呼叫順序相反，類似LIFO(queue)
	//NOTE: 也就是說，這個defer是"先進"
	defer func() {
		fmt.Println("defer 1")
	}()
	//NOTE: 這邊是最後進入(LI)，所以他會最先出去(FO)，然後才是defer1
	defer func() {
		fmt.Println("defer 2")
		num -= 5
		//NOTE: defer中任何變動都會在函數結束執行前生效，並隨著return一起被回傳
	}()

	return
}

func normalReturn() int {
	total := 0
	for total < 10 {
		total += 1
	}

	defer func() {
		//NOTE: 無效，不會影響return total當下的值，仍回傳10
		//NOTE: 因為一般return會建立一個臨時變數去儲存回傳值，有名回傳則不會。
		total -= 5
	}()

	return total
}
func main() {

	  fmt.Println("print namedReturn------", namedReturn())   //15
	fmt.Println("print normalReturn------", normalReturn()) //10

}
