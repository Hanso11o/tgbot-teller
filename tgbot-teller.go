package main

import (
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// create an object through which we will communicate with our bot
var bot *tgbotapi.BotAPI

// create name's bot
var tellerBotName = [2]string{"teller", "tell me"}
var answers = [33]string{
	"Часто самая темная дорога ведет к самому яркому свету.",
	"Никогда не знаешь, что тебя ждет в тени, пока не выйдешь в свет.",
	"Самые страшные решения могут привести к самым важным переменам.",
	"Иногда нужно сделать шаг во мрак, чтобы обрести свое собственное свечение.",
	"Не бойтесь того, что скрывается за ужасом. Там обычно спрятано нечто ценное.",
	"Беспокойтесь не о том, что страшно, а о том, что можно упустить из-за своего страха.",
	"Иногда единственный способ найти ответ — погрузиться в самый мрак, где он скрывается.",
	"Страх — всего лишь отражение того, как сильно вы стремитесь к чему-то.",
	"Тот, кто видел истинный ужас, лучше всех знает цену настоящего счастья.",
	"Свет виден яснее после того, как ты побывал в полной темноте.",
	"Следуй за своим страхом, и он покажет тебе путь к самым неожиданным открытиям.",
	"Иногда тьма нужна, чтобы оценить свет.",
	"То, что кажется страшным, может оказаться всего лишь оболочкой для чего-то прекрасного.",
	"Лишь рискнув погрузиться во тьму, можно обрести свой собственный свет.",
	"Страх — это только знак того, что ты готов выйти за рамки своей зоны комфорта.",
	"Иногда настоящая опасность скрывается там, где ты ее меньше всего ожидаешь.",
	"Только преодолев собственный страх, ты можешь обрести настоящую силу.",
	"Часто то, что страшно, стоит больше всего твоего внимания.",
	"Ужас иногда лишь прелюдия к самым захватывающим приключениям.",
	"Темные времена могут порождать светлые истории.",
	"Не бойся тьмы, ведь именно там могут скрываться самые важные ответы.",
	"Иногда собственный страх — это единственное, что мешает нам открыть что-то действительно ценное.",
	"Истинное сокровище часто скрыто в самом мрачном углу.",
	"Только прошедши через ужас, можно по-настоящему оценить красоту.",
	"Только лишившись света, можно обрести свою собственную звезду.",
	"Самые запутанные дороги приводят к самым интересным открытиям.",
	"Темная сторона жизни неизбежна, но именно там скрывается настоящая суть.",
	"Без тьмы свет не имел бы такой силы.",
	"Часто самые неожиданные моменты бывают самыми важными.",
	"Иногда то, что кажется ужасным, оказывается всего лишь началом самых потрясающих историй.",
	"Да",
	"Нет",
	"Я не псайкер",
}
var chatId int64

// creating a function to connect to a telegram
func connectWithTelegram() {
	var err error
	var token string
	if bot, err = tgbotapi.NewBotAPI(token); err != nil {
		panic("cannot connect to telegram")
	}

}

func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(chatId, msg)
	bot.Send(msgConfig)
}

func isMessageToBot(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false

	}

	msgToLower := strings.ToLower(update.Message.Text)
	for _, name := range tellerBotName {
		if strings.Contains(msgToLower, name) {
			return true
		}
	}
	return false

}

func getAnswer() string {
	index := rand.Intn(len(answers))
	return answers[index]
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatId, getAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	connectWithTelegram()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatId = update.Message.Chat.ID
			sendMessage("Задай вопрос, назвав меня по имени teller или tell me.")
		}
		if isMessageToBot(&update) {
			sendAnswer(&update)

		}

	}
}
