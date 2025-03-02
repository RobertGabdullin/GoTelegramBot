package server

import (
	"github.com/RobertGabdullin/GoTelegramBot/internal/scrapper/handler"
	"net/http"
)

type Server struct {
	handler handler.LinkHandler
	server  *http.Server
	address string
}

func NewServer(linkHandler handler.LinkHandler, address string) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/tg-chat/", linkHandler.TgChat)
	mux.HandleFunc("/links", linkHandler.Links)
	return &Server{
		handler: linkHandler,
		server: &http.Server{
			Addr:    address,
			Handler: mux,
		},
	}
}

func (s *Server) Start() {
	s.server.ListenAndServe()
}
