package main

import (
	"log"
	"os"

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

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text =="/start" 
	}

}
