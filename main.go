package main

import (
	"log"
	"musicbot/client"
	consumer "musicbot/consumer"
)

func main() {
	tgm := client.New("api.telegram.org", "7112110942:AAEjJBRkyhwj9UiIGM8-W3MPkpcDf9cnIP0")
	startConsumer := consumer.New(100, 10, tgm)

	if err := startConsumer.WorkingCycle(); err != nil {
		log.Println(err)
	}

}
