package transport

type LinkUpdateRequest struct {
	Id          int    `json:"id"`
	Url         string `json:"url"`
	Description string `json:"description"`
	TgChatIds   []int  `json:"tg_chat_ids"`
}
