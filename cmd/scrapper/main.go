package main

import (
	"fmt"
	"github.com/RobertGabdullin/GoTelegramBot/configs"
	"github.com/RobertGabdullin/GoTelegramBot/internal/scrapper/handler"
	"github.com/RobertGabdullin/GoTelegramBot/internal/scrapper/server"
	"github.com/RobertGabdullin/GoTelegramBot/internal/scrapper/service"
	"github.com/RobertGabdullin/GoTelegramBot/internal/scrapper/storage"
)

func main() {

	config, err := configs.LoadConfig()
	if err != nil {
		fmt.Println("load config err:", err)
		return
	}

	linkStorage, err := storage.NewPostgresqlLinkTracker(config.Scrapper.DatabaseUrl)
	if err != nil {
		fmt.Println("create storage err:", err)
		return
	}

	linkService := service.NewDBLinkService(linkStorage)
	linkHandler := handler.NewDefaultLinkHandler(linkService)
	linkServer := server.NewServer(linkHandler, config.Scrapper.BaseUrl)
	linkServer.Start()

}
