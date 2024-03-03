package main

import (
	"fmt"
	"time"

	cron "github.com/robfig/cron/v3"
)

func main() {
  c := cron.New()

  c.AddFunc("@every 5s", func() {
    fmt.Println("tick every 1 second")
  })

  c.Start()

  //總執行時間，例如@every5s 
  time.Sleep(time.Duration(10) * time.Second)

  // Stop the Cron job scheduler
  c.Stop()
}