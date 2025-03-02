package bot

import (
	"github.com/RobertGabdullin/GoTelegramBot/configs"
	"github.com/RobertGabdullin/GoTelegramBot/internal/bot/processor"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Bot                  *tgbot.BotAPI
	UpdateChannel        tgbot.UpdatesChannel
	UserMessageProcessor processor.MessageProcessor
}

func NewTelegramBot(config *configs.BotConfig) (*TelegramBot, error) {
	bot, err := tgbot.NewBotAPI(config.Token)
	if err != nil {
		return nil, err
	}

	update := tgbot.NewUpdate(config.UpdateOffset)
	update.Timeout = config.UpdateTimeout
	updateChan := bot.GetUpdatesChan(update)

	messageProcessor := processor.NewTelegramMessageProcessor()

	tgBot := &TelegramBot{
		Bot:                  bot,
		UpdateChannel:        updateChan,
		UserMessageProcessor: messageProcessor,
	}

	return tgBot, nil
}

func (b *TelegramBot) Start() {
	for update := range b.UpdateChannel {
		b.Bot.Send(b.UserMessageProcessor.Process(update))
	}
}
