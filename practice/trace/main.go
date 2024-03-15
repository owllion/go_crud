package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	file, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = trace.Start(file)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main goroutine
	fmt.Println("Hello World!")
}
