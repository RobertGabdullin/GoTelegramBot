package storage

type LinkTrackerStorage interface {
	AddLink(link string) error
	RemoveLink(link string) error
	AddChat(chatId int) error
	RemoveChat(chatId int) error
	AddChatLink(chatId int, linkId int) error
	RemoveChatLink(chatId int, linkId int) error
	GetLinks(chatId int) ([]string, error)
	GetIdByLink(link string) (int, error)
	GetAllLinks() ([]string, error)
	IsChatPresent(chatId int) (bool, error)
	IsLinkPresent(link string) (bool, error)
	IsChatLinkPresent(chatId int, linkId int) (bool, error)
}
