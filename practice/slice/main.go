package main

import "fmt"

func ptrEscape() *int {
	a := 5 //直接分配到heap了，這是編譯器依照逃逸分析做出的判斷
	return &a
}

func sliceEscape() []int {
	n := []int{1, 2, 3, 4} //escape
	return n
}
func mapEscape() map[string]string {
	m_ := map[string]string{ //escape
		"wow": "fuck",
	}
	return m_
}
func intNotEscape() int {
	b := 543
	return b
}
func main() {
	_ = mapEscape() //escape
	_ = intNotEscape()
	resA := ptrEscape() //escape
	b := *resA + 75     //escape
	fmt.Println(b)
	//因為println接受的參數是interface{}(func  Println (a ... interface {}))，編譯器難以確定大小(heap大小不是固定的，比較靈活)，所以永遠都會是分配到heap

}
