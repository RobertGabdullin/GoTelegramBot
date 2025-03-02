package processor

import (
	"github.com/RobertGabdullin/GoTelegramBot/internal/bot/command"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageProcessor interface {
	GetCommandsMap() map[string]command.Command
	Process(update tgbot.Update) tgbot.MessageConfig
}
