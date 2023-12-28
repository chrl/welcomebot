package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strings"
)

var bot *tgbotapi.BotAPI
var message string

func main() {

	s:= os.Getenv("BOT_TOKEN")
	message = os.Getenv("DEFAULT_GREET")
	debug:= os.Getenv("DEBUG")



	if len(message) == 0 {
		message = "Welcome, @%s!"
	}

	var err error

	bot, err = tgbotapi.NewBotAPI(s)
	if err != nil {
		log.Panic(err)
	}

	if len(debug) > 0 {
		bot.Debug = true
	}

	log.Printf("Authorized on account %s\n", bot.Self.UserName)
	log.Println("I'm running!")


	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update:= range updates {

		log.Print("Got update: ")
		log.Printf("%+v", update)

		if update.Message!=nil && update.Message.NewChatMembers != nil {
			handleGroupJoin(update)
		}

		if update.Message!=nil && update.Message.IsCommand() {
			handleCommand(update)
		}
	}

}

func handleGroupJoin(update tgbotapi.Update)  {
	log.Println("Join to the group caught, replying")

	for _, newMember := range update.Message.NewChatMembers {
		// Send a welcome message to the new member
		welcomeMessage := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(message, newMember.UserName))
		bot.Send(welcomeMessage)
	}
}

func handleCommand(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "setwelcome":
		log.Println(update.Message.Text)
		message = strings.Replace(update.Message.Text,"/setwelcome ","",-1)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome message updated")
		bot.Send(msg)
	default:
	}
}