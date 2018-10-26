package main

import (
	"github.com/robfig/cron"
	"log"
	"niurenshuo/models"
	"time"
)

func main() {
	log.Println("Starting...")
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllComment...")
		models.CleanAllComment()
	})

	c.Start()

	timer := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-timer.C:
			timer.Reset(time.Second * 10)
		}
	}
}
