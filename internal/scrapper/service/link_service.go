package service

type LinkService interface {
	Register(chatId int) error
	Track(chatId int, link string) error
	Untrack(chatId int, link string) error
	List(chatId int) ([]string, error)
}
