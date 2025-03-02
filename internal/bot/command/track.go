package command

import (
	"github.com/RobertGabdullin/GoTelegramBot/internal/bot/client"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type Track struct {
	Name        string
	Description string
}

func NewTrack() Track {
	return Track{
		Name:        "/track",
		Description: "Start tracking the link",
	}
}

func (cmd Track) GetName() string {
	return cmd.Name
}

func (cmd Track) GetDescription() string {
	return cmd.Description
}

func (cmd Track) Handle(update tgbot.Update, scrapperClient client.ScrapperClient) tgbot.MessageConfig {
	words := strings.Split(update.Message.Text, " ")
	if len(words) < 2 {
		return tgbot.NewMessage(update.Message.Chat.ID, "No link provided")
	}
	link := words[1]
	message, err := scrapperClient.PostLinks(update.Message.Chat.ID, link)
	if err != nil {
		return tgbot.NewMessage(update.Message.Chat.ID, err.Error())
	}
	return tgbot.NewMessage(update.Message.Chat.ID, message)
}
