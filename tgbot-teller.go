package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// create an object through which we will communicate with our bot
var bot *tgbotapi.BotAPI
// create name's bot
var tellerBotName = [2]string{"teller", "tell me"}


// creating a function to connect to a telegram
func connectWithTelegram() {
	var err error
	var token string
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("cannot connect to telegram")
		_ = token

	}
}

func sendMessage(msg string) {
	// TODO
}

func isMessageToBot(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text  == "" {
		return false

	}

	msgToLower := strings.ToLower(update.Message.Text)
	for _, name := range tellerBotName {
		if strings.Contains(msgToLower, name ) {
			return true
		}
	}
	return false

}

func sendAnswer(update *tgbotapi.Update)

func main() {
	connectWithTelegram()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			sendMessage("Задай вопрос, назвав меня по имени teller или tell me. Ответом на вопрос должны быть \"да\" либо \"Нет\"")
		}
		if isMessageToBot(&update)
		sendAnswer(&update)

	}

}
