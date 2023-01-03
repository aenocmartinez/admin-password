package usecase

import (
	"abix360/src/dao"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type FindPasswordUseCase struct{}

func (useCase *FindPasswordUseCase) Execute(id int64) (dto dto.PasswordDTO, err error) {
	var repository domain.RepositoryPassword = dao.NewPasswordDAO()
	password, err := domain.FindPasswordById(id, repository)
	if err != nil {
		return dto, err
	}

	if !password.Exists() {
		return dto, errors.New("el elemento buscado no existe")
	}

	dto.Id = password.Id()
	dto.Name = password.Name()
	dto.Password = password.Password()
	dto.User = password.User()
	dto.Url = password.Url()

	return dto, err
}
