package command

import (
	"github.com/RobertGabdullin/GoTelegramBot/internal/bot/client"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Start struct {
	Name        string
	Description string
}

func NewStart() Start {
	return Start{
		Name:        "/start",
		Description: "Register user",
	}
}

func (cmd Start) GetName() string {
	return cmd.Name
}

func (cmd Start) GetDescription() string {
	return cmd.Description
}

func (cmd Start) Handle(update tgbot.Update, scrapperClient client.ScrapperClient) tgbot.MessageConfig {
	message, err := scrapperClient.PostTgChat(update.Message.Chat.ID)
	if err != nil {
		return tgbot.NewMessage(update.Message.Chat.ID, err.Error())
	}
	return tgbot.NewMessage(update.Message.Chat.ID, message)
}
