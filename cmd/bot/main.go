package main

import (
	"fmt"
	"github.com/RobertGabdullin/GoTelegramBot/configs"
	"github.com/RobertGabdullin/GoTelegramBot/internal/bot"
)

func main() {

	config, err := configs.LoadConfig()
	if err != nil {
		fmt.Println("Unable to load config: " + err.Error())
		return
	}

	tgBot, err := bot.NewTelegramBot(&config.Bot)
	if err != nil {
		fmt.Println("Unable to create bot: " + err.Error())
		return
	}

	tgBot.Start()

}
