package main

import (
	"fmt"

	cron "github.com/robfig/cron/v3"
)

func main() {
	i := 0
	c := cron.New(cron.WithSeconds())
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})
	c.Start()

	select {}
}
