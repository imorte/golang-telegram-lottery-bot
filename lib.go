package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"time"
	"math/rand"
)

func checkAdminAccess(msg *tgbotapi.Message) bool {
	var info Info

	gdb.Where("id = ?", 1).First(&info)

	if msg.From.UserName != info.Admin {
		print(msg.From.UserName)
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Увы, но вы не можете выполнить эту команду"))
		return false
	} else {
		return true
	}
}

func digitToWord(digit int) string {
	digitToWord := map[int]string {
		1: "один",
		2: "два",
		3: "три",
		4: "четыре",
		5: "пять",
		6: "шесть",
		7: "семь",
		8: "восемь",
		9: "девять",
	}

	return digitToWord[digit]
}

func uniqueRandom(random int, count int) []int {
	var result[] int
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(count)
	for _, r := range p[:random] {
		result = append(result, r + 1)
	}

	return result
}

