package repository

type LinkRepository interface {
	Add(ID int64) error
	Delete(ID int64) error
	GetAll() ([]string, error)
}
