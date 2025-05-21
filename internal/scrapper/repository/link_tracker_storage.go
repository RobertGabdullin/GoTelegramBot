package repository

type LinkTrackerStorage interface {
	AddLink(link string) error
	RemoveLink(link string) error
	AddChat(chatId int64) error
	RemoveChat(chatId int64) error
	AddChatLink(chatId int64, linkId int64) error
	RemoveChatLink(chatId int64, linkId int64) error
	GetLinks(chatId int64) ([]string, error)
	GetIdByLink(link string) (int64, error)
	GetAllLinks() ([]string, error)
	IsChatPresent(chatId int64) (bool, error)
	IsLinkPresent(link string) (bool, error)
	IsChatLinkPresent(chatId int64, linkId int64) (bool, error)
}
