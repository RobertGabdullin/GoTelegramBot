package command

import (
	"github.com/RobertGabdullin/GoTelegramBot/internal/bot/client"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type Untrack struct {
	Name        string
	Description string
}

func NewUntrack() Untrack {
	return Untrack{
		Name:        "/untrack",
		Description: "Stop tracking the link",
	}
}

func (cmd Untrack) GetName() string {
	return cmd.Name
}

func (cmd Untrack) GetDescription() string {
	return cmd.Description
}

func (cmd Untrack) Handle(update tgbot.Update, scrapperClient client.ScrapperClient) tgbot.MessageConfig {
	words := strings.Split(update.Message.Text, " ")
	if len(words) < 2 {
		return tgbot.NewMessage(update.Message.Chat.ID, "No link provided")
	}
	link := words[1]
	message, err := scrapperClient.DeleteLinks(update.Message.Chat.ID, link)
	if err != nil {
		return tgbot.NewMessage(update.Message.Chat.ID, err.Error())
	}
	return tgbot.NewMessage(update.Message.Chat.ID, message)
}
