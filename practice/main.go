package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//NOTE: ... operator
	fmt.Println("...nums--->", justAdd(1, 2, 3, 4, 5)) //15
	//NOTE: pointer
	a := []int64{1, 2, 3, 4, 5}
	getDouble(&a, addOne)
	fmt.Println(a)
	//NOTE: 拼接字串
	//NOTE: 效率排序-> strings.Join ~ string.Builder > bytes.Buffer > "+" > fmt.Sprtinf
	//string.Builder
	var strBuilder strings.Builder
	strBuilder.WriteString("Hello")
	strBuilder.WriteString(" is Builder!")
	fmt.Println("print strings.Builder's result*---", strBuilder.String())

	//byte.Buffer
	var strBuffer bytes.Buffer
	strBuffer.WriteString("Hi")
	strBuffer.WriteString(" is Buffer concatanated string")
	fmt.Println("print bytes.Buffer's result*---", strBuffer.String())

	//strings.Join
	joinStr := strings.Join([]string{"Just", "live", "while", "we're young"}, "-")
	fmt.Println("print strings.Join's result*---", joinStr)
}

func getDouble(nums *[]int64, fn func(num int64) int64) {
	for i := range *nums {
		(*nums)[i] = fn((*nums)[i])
	}

}
func addOne(num int64) int64 {
	return num + 1
}

func justAdd(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}
