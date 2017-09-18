package main

import (
	"gopkg.in/telegram-bot-api.v4"
	//"strings"
	"strings"
	"strconv"
	//"log"
	//"log"
	"log"
)

func reg(msg *tgbotapi.Message, update tgbotapi.Update) {
	//var reply tgbotapi.MessageConfig
	var info Info
	var users User
	var currentUser User
	var textReply string

	gdb.Where("id = ?", 1).First(&info)

	if info.Active == false {
		textReply = "В настоящий момент регистрация закрыта"
	} else {
		gdb.Where("userId = ?", msg.From.ID).First(&currentUser)

		if currentUser.Id > 0 {
			textReply = "Вы уже учавствуете в конкурсе!"
		} else {
			gdb.Create(&users)
			gdb.Model(&users).Update(User{
				UserId: msg.From.ID,
				Username: msg.From.UserName,
				UserNick: msg.From.FirstName + " " + msg.From.LastName,
				IsWinner: false,
			})
			textReply = "Вы зарегестрировались! Если вы выиграете, мы с Вами свяжемся!"
		}
	}

	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, textReply))
}

func start(msg *tgbotapi.Message, update tgbotapi.Update) {
	var reply tgbotapi.MessageConfig
	var seq Sequence
	var users User
	var info Info
	var infoCheck Info
	var textReply string

	gdb.Where("id = ?", 1).First(&infoCheck)

	if infoCheck.Active != true {
		//gdb.Model(&info).Where("id = ?", 1).UpdateColumn("is_active", true)
		gdb.Model(&info).Where("id = ?", 1).Update(&Info{Active: true, IsReady: false})
		gdb.Model(&info).Where("id = ?", 1).UpdateColumn("is_ready", false)
		gdb.Delete(&users)
		gdb.Delete(&seq)
		textReply = "Список очищен и регистрация открыта"
	} else {
		textReply = "Регистрация уже открыта!"
	}

	reply = tgbotapi.NewMessage(msg.Chat.ID, textReply)
	bot.Send(reply)
}

func stop(msg *tgbotapi.Message, update tgbotapi.Update) {
	var reply tgbotapi.MessageConfig
	var textReply string
	var info Info
	var infoCheck Info

	gdb.Where("id = ?", 1).First(&infoCheck)

	if infoCheck.Active == false {
		textReply = "Регистрация уже закрыта!"
	} else {
		gdb.Model(&info).Where("id = ?", 1).UpdateColumn("active", false)
		textReply = "Регистрация закрыта"
	}

	reply = tgbotapi.NewMessage(msg.Chat.ID, textReply)
	bot.Send(reply)
}

func list(msg *tgbotapi.Message) {
	var users []User
	var output string

	gdb.Find(&users)

	if len(users) != 0 {
		for _, i := range users {
			output += "@" + i.Username + "\n"
		}
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, output))
	} else {
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Список пуст"))
	}
}

func startLottery(msg *tgbotapi.Message) {
	var infoCheck Info
	var info Info
	var users []User
	var newUser User
	var output string
	var winners []User
	gdb.Find(&users)
	winnersCount := 3

	gdb.Where("id = ?", 1).First(&infoCheck)

	if infoCheck.IsReady == true {
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Победители уже определены!"))
	} else {
		if len(users) < winnersCount {
			winnersCount = len(users)
		}

		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Определяются победители - " + digitToWord(winnersCount) + "..."))


		for _, i := range uniqueRandom(winnersCount, len(users)) {
			gdb.Model(&newUser).Where("id = ?", i).UpdateColumn("is_winner", true)
		}

		gdb.Model(&info).Where("id = ?", 1).UpdateColumn("is_ready", true)

		output = "Выбраны победители!\n"
		gdb.Where("is_winner = ?", true).Find(&winners)
		for count, i := range winners {
			output += strconv.Itoa(count + 1) + ". " +  i.UserNick + " " + "(@" + i.Username + ")\n"
		}

		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, output))
	}
}

func messageToWinners(msg *tgbotapi.Message) {
	var reply tgbotapi.MessageConfig
	var winnersId []string
	var text string
	message := msg.Text
	cnt := 0

	explodedWinners := strings.Split(message, " ")

	for _, x := range explodedWinners {
		if _, err := strconv.Atoi(x); err != nil {
			if cnt != 0 {
				text += x + " "
			}
			cnt++
		} else {
			winnersId = append(winnersId, x)
		}
	}

	for _, x := range winnersId {
		var winner User
		gdb.Where("id = ?", x).First(&winner)
		reply = tgbotapi.NewMessage(int64(winner.UserId), text)
		_, err := bot.Send(reply)
		if err != nil {
			log.Println(err)
		}
	}

	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Поздравления отправлены!"))
}

func regstop(msg *tgbotapi.Message) {
	var info Info

	gdb.Model(&info).Where("id = ?", 1).UpdateColumn("active", false)
	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Регистрация приостановлена"))
}

func regstart(msg *tgbotapi.Message) {
	var info Info

	gdb.Model(&info).Where("id = ?", 1).UpdateColumn("active", true)
	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Регистрация возобновлена"))
}