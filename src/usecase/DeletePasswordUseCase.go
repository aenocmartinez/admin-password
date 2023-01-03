package usecase

import (
	"abix360/src/dao"
	"abix360/src/domain"
	"errors"
)

type DeletePasswordUseCase struct{}

func (useCase *DeletePasswordUseCase) Execute(id int64) error {
	var repository domain.RepositoryPassword = dao.NewPasswordDAO()

	password, err := domain.FindPasswordById(id, repository)
	if err != nil {
		return err
	}

	if !password.Exists() {
		return errors.New("no existe el elemento")
	}

	password.WithRepository(repository)

	err = password.Delete()
	if err != nil {
		return err
	}

	return err
}
