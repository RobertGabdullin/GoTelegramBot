package handler

import (
	"encoding/json"
	"github.com/RobertGabdullin/GoTelegramBot/internal/scrapper/service"
	"github.com/RobertGabdullin/GoTelegramBot/internal/transport"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

type DefaultLinkHandler struct {
	linkService service.LinkService
}

func NewDefaultLinkHandler(linkService service.LinkService) *DefaultLinkHandler {
	return &DefaultLinkHandler{
		linkService: linkService,
	}
}

func (d DefaultLinkHandler) TgChat(w http.ResponseWriter, r *http.Request) {
	pathPart := strings.TrimPrefix(r.URL.Path, "/tg-chat/")
	id, err := strconv.Atoi(pathPart)

	if err != nil {
		apiError := transport.APIErrorResponse{
			Description:      "Url is not valid",
			Code:             "400",
			ExceptionName:    "Bad Request",
			ExceptionMessage: err.Error(),
		}
		apiErrorJson, _ := json.Marshal(apiError)
		w.WriteHeader(400)
		w.Write(apiErrorJson)
	}

	switch r.Method {
	case http.MethodGet:

	}
}

func (d DefaultLinkHandler) Links(w http.ResponseWriter, r *http.Request) {

}
