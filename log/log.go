package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

func init() {
	log.SetPrefix("【Debug】")
	log.SetFlags(log.Ldate | log.Ltime)
}

func Log(information string, query interface{}, e string) {
	t := time.Now().Local().Format("2006-01-02")

	fileName := fmt.Sprint("./log/", t, ".log")

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
		return
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(information)
	log.Printf("%+v\n", query)
	log.Printf("%+v\n", e)
}

func SetLog(information string, query interface{}) {
	t := time.Now().Local().Format("2006-01-02")

	fileName := fmt.Sprint("./log/", t, ".log")

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(information)
	log.Printf("%+v\n", query)
}