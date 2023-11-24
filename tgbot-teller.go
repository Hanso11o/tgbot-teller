package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// create an object through which we will communicate with our bot
var bot *tgbotapi.BotAPI

func connectWithTelegram() {
	var err error
	var token string
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("cannot connect to telegram")
		_ = token

	}
}

func main() {
	connectWithTelegram()

}
