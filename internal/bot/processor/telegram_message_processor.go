package processor

import (
	"strings"

	"github.com/RobertGabdullin/GoTelegramBot/internal/bot/command"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	helpMessage = "This bot helps you track resource changes on GitHub and StackOverFlow.\n\n" +
		"You can control bot by sending these commands:\n"
)

type TelegramMessageProcessor struct {
	commandMap map[string]command.Command
}

func NewTelegramMessageProcessor() TelegramMessageProcessor {
	cmdMap := map[string]command.Command{
		"/start":   command.NewStart(),
		"/list":    command.NewList(),
		"/track":   command.NewTrack(),
		"/untrack": command.NewUntrack(),
	}
	return TelegramMessageProcessor{
		commandMap: cmdMap,
	}
}

func (processor TelegramMessageProcessor) GetCommandsMap() map[string]command.Command {
	return processor.commandMap
}

func (processor TelegramMessageProcessor) helpHandle() string {
	resultMessage := &strings.Builder{}
	resultMessage.Write([]byte(helpMessage))
	for _, cmd := range processor.commandMap {
		resultMessage.Write([]byte(cmd.GetName()))
		resultMessage.Write([]byte(" - "))
		resultMessage.Write([]byte(cmd.GetDescription()))
		resultMessage.Write([]byte("\n"))
	}
	return resultMessage.String()
}

func (processor TelegramMessageProcessor) Process(update tgbot.Update) tgbot.MessageConfig {
	cmdName := strings.Split(update.Message.Text, " ")[0]
	if cmdName == "/help" {
		return tgbot.NewMessage(update.Message.Chat.ID, processor.helpHandle())
	}
	cmd, exist := processor.commandMap[cmdName]
	if !exist {
		return tgbot.NewMessage(update.Message.Chat.ID, "Command not found."+
			" Type /help to see list of supported commands.")
	}
	return cmd.Handle(update)
}
