package domain

type RepositoryPassword interface {
	Save(password Password) (Password, error)
	Update(password Password) error
	Delete(password Password) error
	FindById(id int64) Password
}
