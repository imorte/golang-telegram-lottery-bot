package main

import (
	"gopkg.in/telegram-bot-api.v4"
)

func reg(msg *tgbotapi.Message, update tgbotapi.Update) {
	//var reply tgbotapi.MessageConfig

	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Hey"))
}