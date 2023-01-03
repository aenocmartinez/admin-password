package domain

type RepositoryPassword interface {
	Save(password Password) (Password, error)
	Update(password Password) (Password, error)
	Delete(password Password) error
	FindById(id int64) (Password, error)
	All() ([]Password, error)
}
