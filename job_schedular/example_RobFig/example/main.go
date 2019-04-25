package main

import (
	"github.com/robfig/cron"
	"fmt"
	"time"
)

func main()  {
	c := cron.New()
	c.AddFunc("* * 1 * * *", func() { fmt.Println("Every sec on the half hour") })
	c.AddFunc("@hourly",      func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()
	// Funcs are invoked in their own goroutine, asynchronously.
	// Funcs may also be added to a running Cron
	c.AddFunc("@daily", func() { fmt.Println("Every day") })
	// Inspect the cron job entries' next and previous run times.
	time.Sleep(5*time.Second)
}