package handler

import "net/http"

type LinkHandler interface {
	TgChat(w http.ResponseWriter, r *http.Request)
	Links(w http.ResponseWriter, r *http.Request)
}
