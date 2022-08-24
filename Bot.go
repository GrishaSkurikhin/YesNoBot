package main

import (
	"log"
	"reflect"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	startMessage = "Для получения ответа отправьте команду /get\nДля выбора фиксированного ответа отправьте команду /yes или /no"
	timeout      = 60
)

func TelegramBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, startMessage)
				bot.Send(msg)

			case "/get":
				sendMessage(bot, &update, false, "")

			case "/yes":
				sendMessage(bot, &update, true, "yes")

			case "/no":
				sendMessage(bot, &update, true, "no")
			}
		}
	}
}

func sendMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update, forced bool, forcedAns string) {
	var response *Response
	var err error
	if forced {
		response, err = GetForcedAnswer(forcedAns)
	} else {
		response, err = GetAnswer()
	}

	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка")
		bot.Send(msg)
	}

	msg := tgbotapi.NewDocumentShare(update.Message.Chat.ID, response.Image)
	msg.Caption = engToRus(response.Answer)
	bot.Send(msg)
}

func engToRus(message string) string {
	switch message {
	case "yes":
		return "✅Да"
	case "no":
		return "❌Нет"
	}
	return ""
}
