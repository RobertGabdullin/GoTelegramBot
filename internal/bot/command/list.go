package command

import (
	"fmt"
	"github.com/RobertGabdullin/GoTelegramBot/internal/bot/client"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type List struct {
	Name        string
	Description string
}

func NewList() List {
	return List{
		Name:        "/list",
		Description: "Display a list of tracked links",
	}
}

func (cmd List) GetName() string {
	return cmd.Name
}

func (cmd List) GetDescription() string {
	return cmd.Description
}

func (cmd List) Handle(update tgbot.Update, scrapperClient client.ScrapperClient) tgbot.MessageConfig {
	links, err := scrapperClient.GetLinks(update.Message.Chat.ID)
	if err != nil {
		return tgbot.NewMessage(update.Message.Chat.ID, err.Error())
	}

	if len(links) == 0 {
		return tgbot.NewMessage(update.Message.Chat.ID, "No links found")
	}
	builder := strings.Builder{}
	builder.WriteString("Here are all links you are tracking:\n")
	for _, link := range links {
		builder.WriteString(fmt.Sprintf("- %s\n", link))
	}
	return tgbot.NewMessage(update.Message.Chat.ID, builder.String())
}
