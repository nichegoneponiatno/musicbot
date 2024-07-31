package main

import (
	"log"
	"musicbot/client"
	consumer "musicbot/consumer"
)

func main() {
	tgm := client.New("api.telegram.org", "token")
	startConsumer := consumer.New(100, 10, tgm)

	if err := startConsumer.WorkingCycle(); err != nil {
		log.Println(err)
	}

}
