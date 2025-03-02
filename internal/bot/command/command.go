package command

import (
	"github.com/RobertGabdullin/GoTelegramBot/internal/bot/client"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Command interface {
	GetName() string
	GetDescription() string
	Handle(update tgbot.Update, scrapperClient client.ScrapperClient) tgbot.MessageConfig
}
