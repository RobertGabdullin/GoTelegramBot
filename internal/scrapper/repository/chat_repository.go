package repository

type ChatRepository interface {
	Add(ID int64) error
	Delete(ID int64) error
}
